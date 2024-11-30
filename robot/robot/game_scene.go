package robot

import (
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (r *RoBot) GetCurSceneInfoScRsp(payloadMsg []byte) {
	msg := decodePayloadToProto(cmd.GetCurSceneInfoScRsp, payloadMsg)
	rsp := msg.(*proto.GetCurSceneInfoScRsp)

	r.Game.EntryId = rsp.Scene.EntryId

	for _, entityGroup := range rsp.Scene.EntityGroupList {
		for _, entity := range entityGroup.EntityList {
			if entity.EntityId == rsp.Scene.LeaderEntityId {
				r.Game.Rot = &Vector{
					X: entity.Motion.Pos.X,
					Y: entity.Motion.Pos.Y,
					Z: entity.Motion.Pos.Z,
				}
				r.Game.Pos = &Vector{
					X: entity.Motion.Rot.X,
					Y: entity.Motion.Rot.Y,
					Z: entity.Motion.Rot.Z,
				}
				break
			}
		}
	}
}

func (r *RoBot) EnterSceneCsReq() {
	if r.KcpConn == nil {
		return
	}
	entryIdList := gdconf.GetEntryIdList()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	entryId := rand.Intn(len(entryIdList)-1) + 1
	rep := &proto.EnterSceneCsReq{
		EntryId: entryIdList[entryId],
	}
	r.send(cmd.EnterSceneCsReq, rep)
}

func (r *RoBot) EnterSceneByServerScNotify(payloadMsg []byte) {
	// r.EnterSceneCsReq()
}
