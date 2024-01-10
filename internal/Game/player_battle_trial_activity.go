package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *Game) StartTrialActivityCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartTrialActivityCsReq, payloadMsg)
	req := msg.(*proto.StartTrialActivityCsReq)

	avatarDemo := gdconf.GetAvatarDemoConfigById(req.AvatarDemoId)
	lineup := g.GetLineUpById(10)

	lineup.AvatarIdList = []uint32{0, 0, 0, 0}
	for id, avatarId := range avatarDemo.TrialAvatarList {
		lineup.AvatarIdList[id] = avatarId
	}

	// g.Send(cmd.SyncServerSceneChangeNotify, req)

	g.SyncLineupNotify(10)

	g.StartTrialEnterSceneByServerScNotify(avatarDemo.MapEntranceID, 0)

	// EnterSceneByServerScNotify

	rsp := &proto.StartTrialActivityScRsp{AvatarDemoId: req.AvatarDemoId}
	g.Send(cmd.StartTrialActivityScRsp, rsp)
}

func (g *Game) StartTrialEnterSceneByServerScNotify(entryId, teleportId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
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
	entityMap := make(map[uint32]*EntityList) // [实体id]怪物群id

	if teleportId != 0 {
		groupID = foorMap.Teleports[teleportId].AnchorGroupID
		anchorID = foorMap.Teleports[teleportId].AnchorID
	} else if anchorID == 0 {
		groupID = foorMap.StartGroupID
		anchorID = foorMap.StartAnchorID
	}

	// 获取队伍
	rsp.Lineup = g.GetLineUpPb(10)

	rsp.Scene = &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,
		LeaderEntityId:     leaderEntityId,
		FloorId:            mapEntrance.FloorID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType),
		PlaneId:            mapEntrance.PlaneID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
	}

	for i := uint32(0); i < 100; i++ {
		rsp.Scene.LightenSectionList = append(rsp.Scene.LightenSectionList, i)
	}

	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 添加队伍角色进实体列表，并设置坐标
	if foorMap.Groups[groupID] == nil {
		return
	}
	for _, anchor := range foorMap.Groups[groupID].AnchorList {
		if anchor.ID == anchorID {
			for id, avatarid := range g.GetLineUpById(10).AvatarIdList {
				if avatarid == 0 {
					continue
				}
				avatarid = gdconf.GetSpecialAvatarById(avatarid).AvatarID
				entityId := uint32(g.GetNextGameObjectGuid())
				entityList := &proto.SceneEntityInfo{
					Actor: &proto.SceneActorInfo{
						AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
						BaseAvatarId: avatarid,
					},
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(anchor.PosX * 1000),
							Y: int32(anchor.PosY * 1000),
							Z: int32(anchor.PosZ * 1000),
						},
						Rot: &proto.Vector{
							X: int32(anchor.RotX * 1000),
							Y: int32(anchor.RotY * 1000),
							Z: int32(anchor.RotZ * 1000),
						},
					},
				}
				// 为进入场景的角色设置与上面相同的实体id
				if id == 0 {
					entityList.EntityId = leaderEntityId
					entityMap[leaderEntityId] = &EntityList{
						Entity:  avatarid,
						GroupId: groupID,
					}
				} else {
					entityList.EntityId = entityId
					entityMap[entityId] = &EntityList{
						Entity:  avatarid,
						GroupId: groupID,
					}
				}
				entityGroup.EntityList = append(entityGroup.EntityList, entityList)
			}
			break
		}
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	// 获取场景实体
	for _, levelGroup := range foorMap.Groups {
		if levelGroup.GroupId == 0 {
			continue
		}
		if len(levelGroup.PropList) == 0 && len(levelGroup.NPCList) == 0 && len(levelGroup.MonsterList) == 0 {
			continue
		}
		rsp.Scene.GroupIdList = append(rsp.Scene.GroupIdList, levelGroup.GroupId)

		// 添加物品实体
		propList := g.GetPropByID(levelGroup, levelGroup.GroupId)
		if len(propList.EntityList) != 0 {
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, propList)
		}
		// 添加怪物实体
		nPCMonsterList, x := g.GetNPCMonsterByID(levelGroup, levelGroup.GroupId, entityMap)
		entityMap = x
		if len(nPCMonsterList.EntityList) != 0 {
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, nPCMonsterList)
		}
		// 添加NPC实体
		nPCList := g.GetNPCByID(levelGroup, levelGroup.GroupId)
		if len(nPCList.EntityList) != 0 {
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, nPCList)
		}
	}

	g.GetBattleState().BattleType = spb.BattleType_Battle_TrialActivity

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}
