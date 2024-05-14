package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
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
	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return
	}

	var groupID = mapEntrance.StartGroupID
	var anchorID = mapEntrance.StartAnchorID

	if foorMap.Teleports[teleportId] != nil {
		groupID = foorMap.Teleports[teleportId].AnchorGroupID
		anchorID = foorMap.Teleports[teleportId].AnchorID
	} else if anchorID == 0 {
		groupID = foorMap.StartGroupID
		anchorID = foorMap.StartAnchorID
	}
	// 获取队伍
	rsp.Lineup = g.GetLineUpPb(g.GetLineUp().MainLineUp)
	curLine := g.GetCurLineUp()
	// 获取世界
	for _, anchor := range foorMap.Groups[groupID].AnchorList {
		if anchor.ID == anchorID {
			curLineUp := g.GetCurLineUp()
			if foorMap.Groups[groupID] == nil || curLineUp == nil {
				return
			}
			pos := &proto.Vector{
				X: int32(anchor.PosX * 1000),
				Y: int32(anchor.PosY * 1000),
				Z: int32(anchor.PosZ * 1000),
			}
			rot := &proto.Vector{
				X: int32(anchor.RotX * 1000),
				Y: int32(anchor.RotY * 1000),
				Z: int32(anchor.RotZ * 1000),
			}
			rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)
			break
		}
	}
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
			groupMap := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
			if groupMap != nil {
				if int(mapEntrance.StartGroupID) > len(groupMap) {
					return
				}

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

				for _, groupInfo := range groupMap {
					mazeGroup := &proto.MazeGroup{GroupId: groupInfo.GroupId}
					mapList.MazeGroupList = append(mapList.MazeGroupList, mazeGroup)
				}

				for _, groupMapList := range groupMap {
					for _, propList := range groupMapList.PropList {
						if propList.AnchorID == 0 {
							continue
						}
						mazeProp := &proto.MazeProp{
							State:    gdconf.CheckPointEnable,
							GroupId:  groupMapList.GroupId,
							ConfigId: propList.ID,
						}
						mapList.MazePropList = append(mapList.MazePropList, mazeProp)
						mapList.UnlockedTeleportList = append(mapList.UnlockedTeleportList, propList.MappingInfoID)
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
