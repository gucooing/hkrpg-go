package player

import (
	"context"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

var LogMsgPlayer uint32 = 2

type GamePlayer struct {
	Uid           uint32
	AccountId     uint32
	GameAppId     uint32
	GateAppId     uint32
	IsJumpMission bool
	Store         *database.GameStore
	IsPE          bool
	// 玩家数据
	PlayerData     *model.PlayerData // 玩家内存
	Platform       spb.PlatformType  // 登录设备
	RouteManager   *RouteManager     // 路由
	SendChan       chan Msg          // 发送消息通道
	RecvChan       chan Msg          // 接收消息通道
	SendCtx        context.Context
	SendCal        context.CancelFunc
	RecvCtx        context.Context
	RecvCal        context.CancelFunc
	IsClosed       bool
	LastUpDataTime int64 // 最近一次的活跃时间
}

type Msg struct {
	CmdId     uint16
	MsgType   MsgType
	PlayerMsg pb.Message
	// command
	CommandList []string
	CommandId   int64
	CommandRsp  string
}

type MsgType int

const (
	Server    MsgType = 1
	Client    MsgType = 2
	GmReq     MsgType = 3
	GmRsp     MsgType = 4
	DailyTask MsgType = 5 // 每日刷新
)

func getCurTime() uint64 {
	return uint64(time.Now().UnixMilli())
}

func (g *GamePlayer) GetPd() *model.PlayerData {
	if g.PlayerData == nil {
		g.PlayerData = model.NewPlayerData()
	}
	return g.PlayerData
}

// 拉取账户数据
func (g *GamePlayer) GetPlayerDateByDb() {
	var err error
	dbPlayer := database.GetPlayerDataByUid(g.Store.PlayerDataMysql,
		g.Store.PeMysql, g.Uid)
	g.PlayerData = new(model.PlayerData)
	if dbPlayer == nil || dbPlayer.BinData == nil {
		dbPlayer = new(constant.PlayerData)
		logger.Info("新账号登录，进入初始化流程")
		g.PlayerData.BasicBin = model.NewBasicBin()
		// 初始化完毕保存账号数据
		dbPlayer.Uid = g.Uid
		dbPlayer.Level = g.GetPd().GetLevel()
		dbPlayer.Exp = g.GetPd().GetMaterialById(model.Exp)
		dbPlayer.Nickname = g.GetPd().GetNickname()
		dbPlayer.BinData, err = pb.Marshal(g.GetPd().BasicBin)
		dbPlayer.DataVersion = g.GetPd().GetDataVersion()
		if err != nil {
			logger.Error("pb marshal error: %v", err)
		}

		if g.IsJumpMission {
			g.FinishAllMission()
			g.FinishAllTutorial()
		}

		err = database.AddPlayerDataByUid(g.Store.PlayerDataMysql,
			g.Store.PeMysql, dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败,err:%s", err.Error())
		}
	} else {
		g.PlayerData.BasicBin = new(spb.PlayerBasicCompBin)
		err = pb.Unmarshal(dbPlayer.BinData, g.PlayerData.BasicBin)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			g.PlayerData.BasicBin = model.NewBasicBin()
		}
	}
	g.PlayerData.BasicBin.Uid = g.Uid
	if g.GetPd().GetIsProficientPlayer() { // 是否是老玩家

	} else {

	}
}

func (g *GamePlayer) UpPlayerDate(status spb.PlayerStatusType) bool {
	var err error
	// 验证状态
	if !g.IsPE {
		redisDb, ok := database.GetPlayerStatus(g.Store.StatusRedis, strconv.Itoa(int(g.Uid)))
		if !ok {
			return false
		}
		statu := new(spb.PlayerStatusRedisData)
		err = pb.Unmarshal(redisDb, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			database.DelPlayerStatus(g.Store.StatusRedis, strconv.Itoa(int(g.Uid)))
			return false
		}
		if statu.GameserverId != g.GameAppId && statu.DataVersion != g.GetPd().GetDataVersion() {
			// 脏数据
			logger.Info("[UID:%v]数据过期，已丢弃", g.Uid)
			return false
		}
	}
	//  确认写入，更新数据版本
	g.GetPd().AddDataVersion()
	dbDate := new(constant.PlayerData)
	dbDate.Uid = g.Uid
	dbDate.Level = g.GetPd().GetLevel()
	dbDate.Exp = g.GetPd().GetMaterialById(model.Exp)
	dbDate.Nickname = g.GetPd().GetNickname()
	dbDate.BinData, err = pb.Marshal(g.GetPd().BasicBin)
	dbDate.DataVersion = g.GetPd().GetDataVersion()
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}
	if err = database.UpPlayerDataByUid(g.Store.PlayerDataMysql,
		g.Store.PeMysql, dbDate); err != nil {
		logger.Error("Update Player error")
		return false
	}
	if !g.SetPlayerPlayerBasicBriefData(status) {
		logger.Error("[UID:%v]玩家简要信息保存失败", g.Uid)
	}
	// 保存地图数据
	for _, block := range g.GetPd().GetAllBlockMap() {
		g.GetPd().UpdateBlock(block)
	}

	return true
}

func (g *GamePlayer) SetPlayerPlayerBasicBriefData(status spb.PlayerStatusType) bool {
	playerBasicBrief := &spb.PlayerBasicBriefData{
		Nickname:          g.GetPd().GetNickname(),
		Level:             g.GetPd().GetLevel(),
		WorldLevel:        g.GetPd().GetWorldLevel(),
		LastLoginTime:     time.Now().Unix(),
		HeadImageAvatarId: g.GetPd().GetHeadIcon(),
		Exp:               g.GetPd().GetMaterialById(model.Exp),
		PlatformType:      g.Platform,
		Uid:               g.Uid,
		Status:            status,
		Signature:         g.GetPd().GetSignature(),
		AssistAvatarList:  g.GetPd().GetAssistAvatarListSpb(g.GetPd().GetAssistAvatarList()),
		DisplayAvatarList: g.GetPd().GetAssistAvatarListSpb(g.GetPd().GetDisplayAvatarlist()),
	}

	bin, err := pb.Marshal(playerBasicBrief)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}
	player := &constant.PlayerBasic{
		Uid:     g.Uid,
		BinData: bin,
	}
	return database.UpdatePlayerBasic(g.Store.PlayerBriefDataRedis,
		g.Store.PeMysql, player)
}

func (g *GamePlayer) Send(cmdId uint16, playerMsg pb.Message) {
	if g.Uid == LogMsgPlayer {
		LogMsgSeed(cmdId, playerMsg)
	}
	g.SendMsg(cmdId, playerMsg)
}

func (g *GamePlayer) DecodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Debug("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	return protoObj
}

var blacklist = []uint16{cmd.SceneEntityMoveScRsp, cmd.SceneEntityMoveCsReq, cmd.PlayerHeartBeatScRsp, cmd.PlayerHeartBeatCsReq} // 黑名单
func IsValid(cmdid uint16) bool {
	for _, value := range blacklist {
		if cmdid == value {
			return false
		}
	}
	return true
}

func LogMsgSeed(cmdId uint16, playerMsg pb.Message) {
	if IsValid(cmdId) {
		data := protojson.Format(playerMsg)
		logger.Debug("S --> C : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
}

func LogMsgRecv(cmdId uint16, payloadMsg pb.Message) {
	if IsValid(cmdId) {
		data := protojson.Format(payloadMsg)
		logger.Debug("C --> S : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
}
