package player

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetRogueScoreRewardInfoCsReq() {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.ScoreRewardInfo = &proto.RogueScoreRewardInfo{
		// TODO 注意时间
		BeginTime:            1706472000,
		EndTime:              1707076799,
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               22,
	}

	g.Send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueTalentInfoCsReq() {
	rsp := &proto.GetRogueTalentInfoScRsp{
		TalentInfo: &proto.RogueTalentInfo{
			RogueTalent: make([]*proto.RogueTalent, 0),
		},
	}

	for _, talent := range gdconf.GetTalentIDList() {
		rogueTalent := &proto.RogueTalent{
			Status: proto.RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK,
			// UnlockProgressList: nil,
			TalentId: talent,
		}
		rsp.TalentInfo.RogueTalent = append(rsp.TalentInfo.RogueTalent, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueInfoCsReq(payloadMsg []byte) {
	beginTime := time.Now().AddDate(0, 0, -1).Unix()
	endTime := beginTime + int64(time.Hour.Seconds()*24*8)
	rsp := new(proto.GetRogueInfoScRsp)
	rogueInfo := &proto.RogueInfo{
		RogueInfoData: &proto.RogueInfoData{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: beginTime,
				EndTime:   4070894399,
				SeasonId:  78,
			},
			RogueScoreInfo: &proto.RogueScoreRewardInfo{
				BeginTime:            beginTime,
				EndTime:              endTime,
				HasTakenInitialScore: true,
				PoolId:               22,
				PoolRefreshed:        true,

				HasTakenReward: nil,
				Score:          999999999,
			},
			RogueAreaInfo: &proto.RogueAreaInfo{RogueArea: make([]*proto.RogueArea, 0)},
			RogueAeonInfo: &proto.RogueAeonInfo{
				// TODO
				AeonIdList:    make([]uint32, 0),
				IsUnlocked:    false,
				UnlockAeonNum: 0,
			},
		},
	}
	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rogueInfo.RogueInfoData.RogueAreaInfo.RogueArea = append(rogueInfo.RogueInfoData.RogueAreaInfo.RogueArea, RogueArea)
	}
	rsp.RogueInfo = rogueInfo

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *GamePlayer) StartRogueCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	req := msg.(*proto.StartRogueCsReq)

	// 更新队伍
	if req.BaseAvatarIdList != nil {
		g.GetLineUpById(9).AvatarIdList = req.BaseAvatarIdList
	}
	g.GetLineUp().MainAvatarId = 0
	// 队伍更新通知
	g.SyncLineupNotify(9)

	entityMap := make(map[uint32]*EntityList) // [实体id]怪物群id
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	// 获取地图
	rogueAreaConfig := gdconf.GetRogueAreaConfigById(strconv.Itoa(int(req.AreaId)))
	if rogueAreaConfig == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	rogueMapID := (rogueAreaConfig.AreaProgress * 100) + rogueAreaConfig.Difficulty
	rogueMapStart := gdconf.GetRogueMapStartById(strconv.Itoa(int(rogueMapID)))
	if rogueMapStart == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	// 获取映射信息
	rogueMap := gdconf.GetRogueRoomIDBySiteID()

	beginTime := time.Now().AddDate(0, 0, -1).Unix()
	endTime := beginTime + int64(time.Hour.Seconds()*24*8)

	// 可独立成单独的方法
	/*
		rogueScoreInfo := &proto.RogueScoreRewardInfo{
			HasTakenInitialScore: true, // 已取得初始积分？
			Score:                0,
			PoolRefreshed:        true, // 刷新？
			TakenScoreRewardList: nil,
			PoolId:               20 + g.PlayerPb.WorldLevel,
		}
	*/
	// 可独立成单独的方法
	roomMap := &proto.RogueMapInfo{
		MapId:     rogueMapID,
		AreaId:    req.AreaId,
		CurSiteId: rogueMapStart.SiteID,
		CurRoomId: rogueMap[1],
		RoomList:  make([]*proto.RogueRoom, 0),
	}
	syncRogueMapRoomScNotify := new(proto.SyncRogueMapRoomScNotify)
	for id, rogue := range rogueMap {
		roomList := &proto.RogueRoom{
			SiteId: id,
			RoomId: rogue,
		}
		if id == rogueMapStart.SiteID {
			roomList.RoomStatus = proto.RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY
			syncRogueMapRoomScNotify.MapId = rogueMapID
			syncRogueMapRoomScNotify.CurRoom = &proto.RogueRoom{
				RoomStatus: proto.RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY,
				SiteId:     id,
				RoomId:     rogue,
			}
		} else {
			roomList.RoomStatus = proto.RogueRoomStatus_ROGUE_ROOM_STATUS_NONE
		}

		roomMap.RoomList = append(roomMap.RoomList, roomList)
	}

	// 区域通知
	g.Send(cmd.SyncRogueMapRoomScNotify, syncRogueMapRoomScNotify)
	// 可独立成单独的方法
	rogueAreaList := make([]*proto.RogueArea, 0)
	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rogueAreaList = append(rogueAreaList, RogueArea)
	}
	// 可独立成单独的方法
	rogueRoom := gdconf.GetRogueRoomById(strconv.Itoa(int(rogueMap[rogueMapStart.SiteID])))
	if rogueRoom == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(rogueRoom.MapEntrance)))
	if mapEntrance == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,
		EntryId:            rogueRoom.MapEntrance,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}

	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}

	// 添加角色信息
	startGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, rogueRoom.GroupID)
	anchor := startGroup.AnchorList[0]
	for id, avatarId := range req.BaseAvatarIdList {
		if avatarId == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: avatarId,
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
				Entity:  avatarId,
				GroupId: rogueRoom.GroupID,
			}
		} else {
			entityList.EntityId = entityId
			entityMap[entityId] = &EntityList{
				Entity:  avatarId,
				GroupId: rogueRoom.GroupID,
			}
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}

	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	// 获取场景实体
	for groupID, _ := range rogueRoom.GroupWithContent {
		sceneGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, stou32(groupID))
		if sceneGroup == nil {
			continue
		}
		scene.GroupIdList = append(scene.GroupIdList, stou32(groupID))

		sceneGroupState := &proto.SceneGroupState{
			GroupId:   stou32(groupID),
			IsDefault: true,
		}

		scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)

		// 添加物品实体
		entityGroupLists := g.GetPropByID(sceneGroup, stou32(groupID))
		// 添加怪物实体
		entityGroupLists, x := g.GetNPCMonsterByID(entityGroupLists, sceneGroup, stou32(groupID), entityMap)
		entityMap = x
		// 添加NPC实体
		entityGroupLists = g.GetNPCByID(entityGroupLists, sceneGroup, stou32(groupID))
		if len(entityGroupLists.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
		}
	}

	rsp := &proto.StartRogueScRsp{
		Scene:  scene,
		Lineup: g.GetLineUpPb(9),
		RogueInfo: &proto.RogueInfo{
			RogueInfoData: &proto.RogueInfoData{
				RogueSeasonInfo: &proto.RogueSeasonInfo{
					BeginTime: beginTime,
					EndTime:   4070894399,
					SeasonId:  78,
				},
				RogueScoreInfo: &proto.RogueScoreRewardInfo{
					BeginTime:            beginTime,
					EndTime:              endTime,
					HasTakenInitialScore: true,
					PoolId:               22,
					PoolRefreshed:        true,

					HasTakenReward: nil,
					Score:          999999999,
				},
				RogueAreaInfo: &proto.RogueAreaInfo{RogueArea: make([]*proto.RogueArea, 0)},
				RogueAeonInfo: &proto.RogueAeonInfo{
					// TODO
					AeonIdList:    make([]uint32, 0),
					IsUnlocked:    false,
					UnlockAeonNum: 0,
				},
			},
			RogueCurrentInfo: &proto.RogueCurrentInfo{
				RogueAvatarInfo: &proto.RogueAvatarInfo{
					BaseAvatarIdList: req.BaseAvatarIdList,
				},
				RoomMap: roomMap,
			},
		},
	}

	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rsp.RogueInfo.RogueInfoData.RogueAreaInfo.RogueArea = append(rsp.RogueInfo.RogueInfoData.RogueAreaInfo.RogueArea, RogueArea)
	}

	g.Player.EntityList = entityMap
	g.GetBattleState().BattleType = spb.BattleType_Battle_ROGUE
	g.Send(cmd.StartRogueScRsp, rsp)
}
