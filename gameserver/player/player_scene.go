package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 通知客户端进入场景
func (g *GamePlayer) EnterSceneByServerScNotify(entryId, teleportId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
	if mapEntrance == nil {
		return
	}
	teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if teleportsMap == nil {
		return
	}

	var anchorID = mapEntrance.StartAnchorID
	var groupID = mapEntrance.StartGroupID
	var pos *proto.Vector
	var rot *proto.Vector
	if teleportsMap[teleportId] == nil {
		teleportId = 0
	}
	// 获取队伍
	rsp.Lineup = g.GetLineUpPb(g.GetLineUp().MainLineUp)
	curLine := g.GetCurLineUp()
	// 获取世界
	if teleportId != 0 {
	te:
		for _, teleports := range teleportsMap {
			for id, teleport := range teleports.Teleports {
				if id == teleportId {
					pos = &proto.Vector{
						X: int32(teleport.PosX * 1000),
						Y: int32(teleport.PosY * 1000),
						Z: int32(teleport.PosZ * 1000),
					}
					rot = &proto.Vector{
						X: int32(teleport.RotX * 1000),
						Y: int32(teleport.RotY * 1000),
						Z: int32(teleport.RotZ * 1000),
					}
					break te
				}
			}
		}
	} else {
		if teleportsMap[groupID] != nil {
			for _, anchor := range teleportsMap[groupID].Teleports {
				if anchor.AnchorID == anchorID {
					pos = &proto.Vector{
						X: int32(anchor.PosX * 1000),
						Y: int32(anchor.PosY * 1000),
						Z: int32(anchor.PosZ * 1000),
					}
					rot = &proto.Vector{
						X: int32(anchor.RotX * 1000),
						Y: int32(anchor.RotY * 1000),
						Z: int32(anchor.RotZ * 1000),
					}
					break
				}
			}
		}
	}

	rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

// 传送到指定位置
func (g *GamePlayer) SceneByServerScNotify(entryId uint32, pos, rot *proto.Vector) {
	rsp := new(proto.EnterSceneByServerScNotify)
	// 获取队伍
	rsp.Lineup = g.GetLineUpPb(g.GetLineUp().MainLineUp)
	curLine := g.GetCurLineUp()
	rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) HandleGetEnteredSceneCsReq(payloadMsg []byte) {
	rsp := new(proto.GetEnteredSceneScRsp)
	db := g.GetScene()
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(db.EntryId)))
	if mapEntrance == nil {
		return
	}
	enteredSceneInfo := &proto.EnteredSceneInfo{
		FloorId: mapEntrance.FloorID,
		PlaneId: mapEntrance.PlaneID,
	}
	rsp.EnteredSceneInfo = []*proto.EnteredSceneInfo{enteredSceneInfo}

	g.Send(cmd.GetEnteredSceneScRsp, rsp)
}

// 客户端登录需要的包，不是传送的通知包
func (g *GamePlayer) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	pos := g.GetPosPb()
	rot := g.GetRotPb()
	dbScene := g.GetScene()
	curLine := g.GetCurLineUp()

	rsp := new(proto.GetCurSceneInfoScRsp)
	rsp.Scene = g.GetSceneInfo(dbScene.EntryId, pos, rot, curLine)
	g.Send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *GamePlayer) HanldeGetSceneMapInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetSceneMapInfoCsReq, payloadMsg)
	req := msg.(*proto.GetSceneMapInfoCsReq)

	// 1000101 1000001

	rsp := new(proto.GetSceneMapInfoScRsp)
	for _, entryId := range req.EntryIdList {
		mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
		if mapEntrance != nil {
			teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
			if teleportsMap != nil {
				mapList := &proto.MazeMapData{
					LightenSectionList: make([]uint32, 0),
					UnlockedChestList: []*proto.MazeChest{
						{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_NORMAL},
						{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
						{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
					},
					UnlockedTeleportList: make([]uint32, 0),
				}

				mapList.EntryId = entryId

				for i := uint32(0); i < 100; i++ {
					mapList.LightenSectionList = append(mapList.LightenSectionList, i)
				}

				for _, teleports := range teleportsMap {
					mazeGroup := &proto.MazeGroup{GroupId: teleports.GroupId}
					mapList.MazeGroupList = append(mapList.MazeGroupList, mazeGroup)
				}

				for _, groupMapList := range teleportsMap {
					for _, teleports := range groupMapList.Teleports {
						mazeProp := &proto.MazeProp{
							State:    gdconf.GetStateValue("CheckPointEnable"), // 默认解锁
							GroupId:  groupMapList.GroupId,
							ConfigId: teleports.ID,
						}
						mapList.MazePropList = append(mapList.MazePropList, mazeProp)
						mapList.UnlockedTeleportList = append(mapList.UnlockedTeleportList, teleports.MappingInfoID)
					}
				}
				rsp.MapList = append(rsp.MapList, mapList)
			}
		}
	}

	g.Send(cmd.GetSceneMapInfoScRsp, rsp)
}

func (g *GamePlayer) EnterSceneCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterSceneCsReq, payloadMsg)
	req := msg.(*proto.EnterSceneCsReq)
	rsp := &proto.GetEnteredSceneScRsp{}

	g.EnterSceneByServerScNotify(req.EntryId, req.TeleportId)
	g.SetCurEntryId(req.EntryId)

	g.Send(cmd.EnterSceneScRsp, rsp)
	g.Send(cmd.SceneUpdatePositionVersionNotify, rsp)
}

func (g *GamePlayer) InteractPropCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.InteractPropCsReq, payloadMsg)
	req := msg.(*proto.InteractPropCsReq)

	rsp := new(proto.InteractPropScRsp)
	rsp.PropEntityId = req.PropEntityId

	g.Send(cmd.InteractPropScRsp, rsp)
}
