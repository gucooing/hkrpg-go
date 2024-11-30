package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

// 获取模拟宇宙信息
func GetRogueInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetRogueInfoScRsp)
	rsp.RogueGameInfo = g.GetPd().GetQuestRogueInfo()

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func GetRogueScoreRewardInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetRogueScoreRewardInfoScRsp{
		Retcode: 0,
		Info:    g.GetPd().GetQuestRogueScoreRewardInfo(),
	}

	g.Send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func GetRogueInitialScoreCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetRogueInitialScoreScRsp{
		RogueScoreRewardInfo: g.GetPd().GetQuestRogueScoreRewardInfo(),
		Retcode:              0,
	}

	g.Send(cmd.GetRogueInitialScoreScRsp, rsp)
}

// 技能树
func GetRogueTalentInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetRogueTalentInfoScRsp{
		TalentInfoList: &proto.RogueTalentInfoList{
			TalentInfo: make([]*proto.RogueTalentInfo, 0),
		},
	}

	for _, talent := range gdconf.GetTalentIDList() {
		rogueTalent := &proto.RogueTalentInfo{
			Status:   proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE,
			TalentId: talent,
		}
		rsp.TalentInfoList.TalentInfo = append(rsp.TalentInfoList.TalentInfo, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func TakeRogueScoreRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeRogueScoreRewardCsReq)
	addItem := model.NewAddItem(nil)
	rsp := &proto.TakeRogueScoreRewardScRsp{
		PoolId:               req.PoolId,
		Retcode:              0,
		Reward:               &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		RogueScoreRewardInfo: nil,
	}
	year, week := time.Now().ISOWeek()
	db := g.GetPd().GetQuestRogueHistoryById(uint32((year%10+(year/10)%10*10)*100 + week))
	if db.RowInfo == nil {
		db.RowInfo = make(map[uint32]bool)
	}
	for _, rowId := range req.RowList {
		conf := gdconf.GetRogueScoreReward(req.PoolId, rowId)
		if conf == nil {
			continue
		}
		if !db.RowInfo[rowId] {
			addItem.PileItem = append(addItem.PileItem, model.GetRewardData(conf.Reward)...)
			db.RowInfo[rowId] = true
		}
	}
	rsp.RogueScoreRewardInfo = g.GetPd().GetQuestRogueScoreRewardInfo()
	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	rsp.Reward.ItemList = addItem.ItemList
	g.Send(cmd.TakeRogueScoreRewardScRsp, rsp)
}

// 开始模拟宇宙
func StartRogueCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartRogueCsReq)
	rsp := &proto.StartRogueScRsp{}
	defer g.Send(cmd.StartRogueScRsp, rsp)
	curRogue, err := g.GetPd().NewQuestRogue(req)
	if err != proto.Retcode(0) {
		rsp.Retcode = uint32(err)
		return
	}
	if true { // 如果是上一次通关过就添加初始奖励
		g.GetPd().AddRogueActionBonusSelect([]uint32{4, 5, 6})
	}
	// 更新队伍
	g.SetBattleLineUp(model.Rogue, req.BaseAvatarIdList)
	// 准备工作就绪,告知客户端
	g.SyncRogueStatusScNotify()

	rsp.Lineup = g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Rogue))

	rsp.Scene = g.GetPd().GetRogueScene(curRogue.GetQuestRogue().RogueRoomMap[curRogue.GetQuestRogue().CurSiteId].RoomId)
	rsp.RogueGameInfo = g.GetPd().GetQuestRogueInfo()

	g.RogueAction() // 模拟宇宙自动化

	g.Send(cmd.StartRogueScRsp, rsp)
}
