package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 通知客户端进入场景
func (g *Game) EnterSceneByServerScNotify(entryId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)

	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))

	groupMap := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if int(mapEntrance.StartGroupID) > len(groupMap) {
		g.send(cmd.EnterSceneByServerScNotify, rsp)
		return
	}
	group := groupMap[mapEntrance.StartGroupID]

	rsp.Scene = &proto.SceneInfo{
		WorldId:         101,
		LeaderEntityId:  1,
		FloorId:         mapEntrance.FloorID,
		GameModeType:    2,
		PlaneId:         mapEntrance.PlaneID,
		EntryId:         entryId,
		EntityGroupList: make([]*proto.SceneEntityGroupInfo, 0),
	}
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, anchor := range group.AnchorList {
		if anchor.ID == mapEntrance.StartAnchorID {
			for _, avatarid := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
				if avatarid == 0 {
					continue
				}
				entityList := &proto.SceneEntityInfo{
					EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
						AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
						BaseAvatarId: avatarid,
					}},
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(anchor.PosX * 1000),
							Y: int32(anchor.PosY * 1000),
							Z: int32(anchor.PosZ * 1000),
						},
						Rot: &proto.Vector{
							X: int32(anchor.RotX),
							Y: int32(anchor.RotY),
							Z: int32(anchor.RotZ),
						},
					},
					EntityId: uint32(g.GetNextGameObjectGuid()),
				}
				entityGroup.EntityList = append(entityGroup.EntityList, entityList)
			}
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

			// TODO 数据保存问题？
			g.Player.Pos = &Vector{
				X: int(anchor.PosX * 1000),
				Y: int(anchor.PosY * 1000),
				Z: int(anchor.PosZ * 1000),
			}
			g.Player.Rot = &Vector{
				X: int(anchor.RotX),
				Y: int(anchor.RotY),
				Z: int(anchor.RotZ),
			}
			g.Player.DbScene.EntryId = entryId
			break
		}
	}

	// 获取场景实体
	for _, levelGroup := range gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID) {
		entityGroupList := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		for _, propList := range levelGroup.PropList {
			entityList := &proto.SceneEntityInfo{
				GroupId:  levelGroup.GroupId,
				InstId:   propList.ID,
				EntityId: uint32(g.GetNextGameObjectGuid()),
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: int32(propList.PosX * 1000),
						Y: int32(propList.PosY * 1000),
						Z: int32(propList.PosZ * 1000),
					},
					Rot: &proto.Vector{
						X: 0,
						Y: 0,
						Z: 0,
					},
				},
				EntityCase: &proto.SceneEntityInfo_Prop{Prop: &proto.ScenePropInfo{
					PropId:    propList.PropID,
					PropState: 0,
				}},
			}
			entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
		}
		rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupList)
	}

	g.send(cmd.EnterSceneByServerScNotify, rsp)

	g.UpDataPlayer()
}

func (g *Game) GetRogueScoreRewardInfoCsReq() {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.ScoreRewardInfo = &proto.RogueScoreRewardInfo{
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               22,
	}

	g.send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *Game) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(g.Player.DbScene.EntryId)))

	rsp := new(proto.GetCurSceneInfoScRsp)
	pos := g.Player.Pos
	rot := g.Player.Rot
	rsp.Scene = &proto.SceneInfo{
		WorldId:         g.Player.DbScene.WorldId,
		LeaderEntityId:  1,
		FloorId:         mapEntrance.FloorID,
		GameModeType:    2,
		PlaneId:         mapEntrance.PlaneID,
		EntryId:         g.Player.DbScene.EntryId,
		EntityGroupList: make([]*proto.SceneEntityGroupInfo, 0),
	}
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, avatarid := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
		if avatarid == 0 {
			continue
		}
		entityList := &proto.SceneEntityInfo{
			EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: avatarid,
			}},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(pos.X),
					Y: int32(pos.Y),
					Z: int32(pos.Z),
				},
				Rot: &proto.Vector{
					X: int32(rot.X),
					Y: int32(rot.Y),
					Z: int32(rot.Z),
				},
			},
			EntityId: uint32(g.GetNextGameObjectGuid()),
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	// 获取场景实体
	for _, levelGroup := range gdconf.GetGroupById(20001, 20001001) {
		entityGroupList := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		for _, propList := range levelGroup.PropList {
			entityList := &proto.SceneEntityInfo{
				GroupId:  levelGroup.GroupId,
				InstId:   propList.ID,
				EntityId: uint32(g.GetNextGameObjectGuid()),
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: int32(propList.PosX * 1000),
						Y: int32(propList.PosY * 1000),
						Z: int32(propList.PosZ * 1000),
					},
					Rot: &proto.Vector{
						X: 0,
						Y: 0,
						Z: 0,
					},
				},
				EntityCase: &proto.SceneEntityInfo_Prop{Prop: &proto.ScenePropInfo{
					PropId:    propList.PropID,
					PropState: 0,
				}},
			}
			entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
		}
		rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupList)
	}

	g.send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *Game) HanldeGetSceneMapInfoCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.GetSceneMapInfoCsReq, payloadMsg)
	req := msg.(*proto.GetSceneMapInfoCsReq)

	entryId := req.EntryIdList[0]

	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))

	if mapEntrance != nil {
		groupMap := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
		if groupMap != nil {
			if int(mapEntrance.StartGroupID) > len(groupMap) {
				return
			}

			rsp := &proto.GetSceneMapInfoScRsp{
				LightenSectionList: make([]uint32, 0),
				UnlockedChestList: []*proto.MazeChest{
					{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_NORMAL},
					{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
					{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
				},
			}

			rsp.EntryId = entryId

			for i := uint32(0); i < 100; i++ {
				rsp.LightenSectionList = append(rsp.LightenSectionList, i)
			}

			for _, groupInfo := range groupMap {
				mazeGroup := &proto.MazeGroup{GroupId: groupInfo.GroupId}
				rsp.MazeGroupList = append(rsp.MazeGroupList, mazeGroup)
			}

			for _, groupMapList := range groupMap {
				for _, propList := range groupMapList.PropList {
					mazeProp := &proto.MazeProp{
						State:    8,
						GroupId:  propList.AnchorGroupID,
						ConfigId: propList.ID,
					}
					rsp.MazePropList = append(rsp.MazePropList, mazeProp)
				}
			}

			g.send(cmd.GetSceneMapInfoScRsp, rsp)
		}
	}
}

func (g *Game) EnterSceneCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.EnterSceneCsReq, payloadMsg)
	req := msg.(*proto.EnterSceneCsReq)

	g.EnterSceneByServerScNotify(req.EntryId)

	rsp := &proto.GetEnteredSceneScRsp{}

	g.send(cmd.EnterSceneScRsp, rsp)
}
