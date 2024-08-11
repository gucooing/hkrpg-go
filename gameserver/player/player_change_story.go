package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) StoryLineInfoScNotify() {
	notify := &proto.StoryLineInfoScNotify{
		CurStoryLineId:            0,
		TrialAvatarIdList:         make([]uint32, 0),
		UnfinishedStoryLineIdList: []uint32{0},
	}

	if db := g.GetCurChangeStoryInfo(); db != nil {
		conf := gdconf.GetStroyLineTrialAvatarData(db.ChangeStoryId)
		if conf != nil {
			notify.CurStoryLineId = db.ChangeStoryId
			notify.TrialAvatarIdList = conf.TrialAvatarList
		}
	}

	for _, changeStory := range g.GetAllChangeStoryInfo() {
		notify.UnfinishedStoryLineIdList = append(notify.UnfinishedStoryLineIdList, changeStory.ChangeStoryId)
	}

	g.Send(cmd.StoryLineInfoScNotify, notify)
}

func (g *GamePlayer) ChangeStoryLineFinishScNotify() {
	notify := &proto.ChangeStoryLineFinishScNotify{
		CurStoryLineId: 0,
		Action:         0,
	}
	if db := g.GetCurChangeStoryInfo(); db != nil {
		notify.CurStoryLineId = db.ChangeStoryId
		notify.Action = proto.ChangeStoryLineAction_ChangeStoryLineAction_Client
	} else {
		return
	}
	g.Send(cmd.ChangeStoryLineFinishScNotify, notify)
}
