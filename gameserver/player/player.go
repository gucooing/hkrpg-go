package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

var LogMsgPlayer uint32 = 0
var ISPE = false

type GamePlayer struct {
	Uid uint32
	// 玩家数据
	PlayerData     *model.PlayerData // 玩家内存
	Platform       spb.PlatformType  // 登录设备
	LoginRandom    uint64
	RouteManager   *RouteManager // 路由
	SendChan       chan Msg      // 发送消息通道
	RecvChan       chan Msg      // 接收消息通道
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

func NewPlayer(uid uint32) *GamePlayer {
	g := new(GamePlayer)
	g.Uid = uid
	g.RecvChan = make(chan Msg, 100)
	g.SendChan = make(chan Msg, 100)
	g.RouteManager = NewRouteManager(g)
	g.LastUpDataTime = time.Now().Unix()
	g.GetPlayerDateByDb()                         // 拉取数据
	g.LoginRandom = random.GetTimeRand().Uint64() // 设置此次随机

	return g
}

// 拉取账户数据
func (g *GamePlayer) GetPlayerDateByDb() {
	var err error
	dbPlayer := database.GetPlayerDataByUid(database.GSS.PlayerDataMysql, g.Uid)
	if g.PlayerData != nil { // 如果有数据了就不需要去拉取
		return
	}
	g.PlayerData = new(model.PlayerData)
	if dbPlayer == nil || dbPlayer.BinData == nil {
		dbPlayer = new(constant.PlayerData)
		logger.Info("新账号登录，进入初始化流程")
		g.PlayerData = model.NewPlayerData()
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

		err = database.AddPlayerDataByUid(database.GSS.PlayerDataMysql, dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败,err:%s", err.Error())
		}
	} else {
		g.PlayerData.BasicBin = new(spb.PlayerBasicCompBin)
		err = pb.Unmarshal(dbPlayer.BinData, g.PlayerData.BasicBin)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			g.PlayerData = model.NewPlayerData()
		}
	}
	if g.GetPd().GetBasicBin().IsJumpMission {
		g.FinishAllTutorial()
	}
	g.PlayerData.BasicBin.Uid = g.Uid
}

func (g *GamePlayer) UpPlayerDate(status spb.PlayerStatusType) bool {
	var err error
	// 验证状态
	if !ISPE {
		redisDb, ok := database.GetPlayerStatus(database.GSS.StatusRedis, g.Uid)
		if !ok {
			return false
		}
		statu := new(spb.PlayerStatusRedisData)
		err = pb.Unmarshal(redisDb, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			database.DelPlayerStatus(database.GSS.StatusRedis, g.Uid)
			return false
		}
		if statu.LoginRand != g.LoginRandom && statu.DataVersion != g.GetPd().GetDataVersion() {
			// 脏数据
			logger.Info("[UID:%v]数据过期，已丢弃", g.Uid)
			return false
		}
		if status == spb.PlayerStatusType_PLAYER_STATUS_OFFLINE {
			database.DelPlayerStatus(database.GSS.StatusRedis, g.Uid)
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
	if err = database.UpPlayerDataByUid(database.GSS.PlayerDataMysql, dbDate); err != nil {
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
		ChatBubbleId:      g.GetPd().GetPhoneData().CurChatBubble,
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
	return database.UpdatePlayerBasic(database.GSS.PlayerBriefDataRedis,
		database.PE, player)
}

func (g *GamePlayer) Send(cmdId uint16, playerMsg pb.Message) {
	if g.Uid == LogMsgPlayer {
		LogMsgSeed(cmdId, playerMsg)
	}
	g.SendMsg(cmdId, playerMsg)
}

func (g *GamePlayer) playerKickOutScNotify() {
	g.Send(cmd.PlayerKickOutScNotify, &proto.PlayerKickOutScNotify{
		KickType:  proto.KickType_KICK_BY_GM,
		BlackInfo: &proto.BlackInfo{},
	})
}

func (g *GamePlayer) Close() {
	if g.IsClosed {
		return
	}
	g.IsClosed = true
	// 等待所有待发送的消息发送完毕
	for {
		if len(g.SendChan) == 0 {
			time.Sleep(time.Millisecond * 100)
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	logger.Info("[UID:%v]玩家下线GAME", g.Uid)
	close(g.RecvChan)
	close(g.SendChan)
	// 保存数据
	g.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_OFFLINE)
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
