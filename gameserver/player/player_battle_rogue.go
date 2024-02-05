package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

// 区域通知
func (g *GamePlayer) SyncRogueMapRoomScNotify() {
	rogue := g.GetDbRogue()

	notify := &proto.SyncRogueMapRoomScNotify{
		CurRoom: &proto.RogueRoom{
			RoomStatus: proto.RogueRoomStatus(rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomStatus),
			SiteId:     rogue.CurRogue.CurSiteId,
			RoomId:     rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomId,
		},
		MapId: rogue.CurRogue.RogueMapID,
	}
	g.Send(cmd.SyncRogueMapRoomScNotify, notify)
}

func (g *GamePlayer) SyncRogueVirtualItemInfoScNotify() {
	notify := &proto.SyncRogueVirtualItemInfoScNotify{
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
			Money: g.GetDbRogue().CurRogue.CosmicFragment,
			X:     8,
		},
	}

	g.Send(cmd.SyncRogueVirtualItemInfoScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonPendingActionScNotify(buffIdList []uint32) {
	rogue := g.GetDbRogue()
	notify := &proto.SyncRogueCommonPendingActionScNotify{
		RogueCommonPendingAction: &proto.RogueCommonPendingAction{
			Num: uint32(len(buffIdList)),
			RogueAction: &proto.RogueAction{
				BuffSelectInfo: &proto.RogueCommonBuffSelectInfo{
					HandbookUnlockBuffIdList: buffIdList,
					CanRoll:                  true,
					MazeBuffList:             make([]*proto.RogueCommonBuff, 0),
					SelectBuffSourceHint:     1,
					SourceCurCount:           rogue.CurRogue.CurSiteId,
					SourceTotalCount:         1,
				},
			},
		},
		MapId: rogue.CurRogue.RogueMapID,
	}

	for _, buffId := range buffIdList {
		rogueCommonBuff := &proto.RogueCommonBuff{
			Level:  1,
			BuffId: buffId,
		}
		notify.RogueCommonPendingAction.RogueAction.BuffSelectInfo.MazeBuffList = append(notify.RogueCommonPendingAction.RogueAction.BuffSelectInfo.MazeBuffList, rogueCommonBuff)
	}

	g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
}

func (g *GamePlayer) SyncEntityBuffChangeListScNotify(buffIdList []uint32) {
	rogue := g.GetRogue()
	notify := &proto.SyncEntityBuffChangeListScNotify{
		EntityBuffInfoList: make([]*proto.EntityBuffChangeInfo, 0),
	}
	if len(buffIdList) == 0 {
		for id, buff := range rogue.BuffList {
			rntityBuffChangeInfo := &proto.EntityBuffChangeInfo{
				AddBuffInfo: &proto.BuffInfo{
					BuffId:    id,
					Level:     buff.Level,
					AddTimeMs: buff.AddTimeMs,
					LifeTime:  -1,
					Count:     4294967295,
				},
				RemoveBuffId: 0,
				EntityId:     6291457,
			}
			notify.EntityBuffInfoList = append(notify.EntityBuffInfoList, rntityBuffChangeInfo)
		}
	} else {
		for _, id := range buffIdList {
			if rogue.BuffList[id] == nil {
				continue
			}
			rntityBuffChangeInfo := &proto.EntityBuffChangeInfo{
				AddBuffInfo: &proto.BuffInfo{
					BuffId:    id,
					Level:     rogue.BuffList[id].Level,
					AddTimeMs: rogue.BuffList[id].AddTimeMs,
					LifeTime:  -1,
					Count:     4294967295,
				},
				RemoveBuffId: 0,
				EntityId:     6291457,
			}
			notify.EntityBuffInfoList = append(notify.EntityBuffInfoList, rntityBuffChangeInfo)
		}
	}

	g.Send(cmd.SyncEntityBuffChangeListScNotify, notify)
}

func (g *GamePlayer) CommonRogueUpdateScNotify() {
	rogue := g.GetDbRogue()
	notify := &proto.CommonRogueUpdateScNotify{
		RogueUpdate: &proto.RogueUpdate{
			MapId:  rogue.CurRogue.RogueMapID,
			AreaId: rogue.CurRogue.CurAreaId,
		},
	}

	g.Send(cmd.CommonRogueUpdateScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonActionResultScNotify(buffId uint32) {
	rogue := g.GetRogue()
	if rogue.BuffList[buffId] == nil {
		return
	}
	notify := &proto.SyncRogueCommonActionResultScNotify{
		Action: &proto.RogueActionResult{
			ActionData: &proto.RogueActionResultData{
				AddBuffList: &proto.RogueBuffData{
					Level:  rogue.BuffList[buffId].Level,
					BuffId: buffId,
				},
			},
			Source: proto.RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_SELECT,
		},
		MapId: g.GetDbRogue().CurRogue.RogueMapID,
	}

	g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}

func (g *GamePlayer) GetRogueScoreRewardInfoCsReq() {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.ScoreRewardInfo = &proto.RogueScoreRewardInfo{
		// TODO 注意时间
		BeginTime:            1706472000,
		EndTime:              4294967295,
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               22,
		Score:                5000, // 本期分数
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
			Status:   proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE,
			TalentId: talent,
		}
		rsp.TalentInfo.RogueTalent = append(rsp.TalentInfo.RogueTalent, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueHandbookDataScRsp() {
	rsp := &proto.GetRogueHandbookDataScRsp{
		HandbookInfo: &proto.RogueHandbookData{
			MiracleList: nil,
			RogueEvent:  nil,
			BuffList:    make([]*proto.RogueHandbookBuff, 0),
		},
	}

	// 解锁全部祝福
	allBuffList := gdconf.GetAllBuff()
	for _, id := range allBuffList {
		buff := &proto.RogueHandbookBuff{
			BuffId: id,
		}
		rsp.HandbookInfo.BuffList = append(rsp.HandbookInfo.BuffList, buff)
	}

	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *GamePlayer) GetRogueInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueInfoScRsp)
	rsp.RogueInfo = g.GetRogueInfo()

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *GamePlayer) StartRogueCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	req := msg.(*proto.StartRogueCsReq)

	if req.BaseAvatarIdList == nil || req.AreaId == 0 {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	g.NewRogue(req.BaseAvatarIdList, req.AreaId)

	//
	g.SyncRogueVirtualItemInfoScNotify()
	syncRogueStatusScNotify := &proto.SyncRogueStatusScNotify{
		RogueStatus: proto.RogueStatus_ROGUE_STATUS_DOING,
	}
	g.Send(cmd.SyncRogueStatusScNotify, syncRogueStatusScNotify)
	// 区域通知
	g.SyncRogueMapRoomScNotify()
	g.Send(cmd.SyncServerSceneChangeNotify, nil)
	// 队伍更新通知
	g.SyncLineupNotify(9)
	var buffList []uint32
	g.SyncEntityBuffChangeListScNotify(buffList)
	g.CommonRogueUpdateScNotify()
	rogue := g.GetDbRogue()

	scene, avatarEntity, monsterEntity := g.GetRogueScene(rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomId)
	if scene == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}

	g.GetSceneEntity().AvatarEntity = avatarEntity
	g.GetSceneEntity().MonsterEntity = monsterEntity

	rsp := &proto.StartRogueScRsp{
		Scene:     scene,
		Lineup:    g.GetLineUpPb(9),
		RogueInfo: g.GetRogueInfo(),
	}
	rsp.RogueInfo.RogueCurrentInfo = &proto.RogueCurrentInfo{
		PendingAction: &proto.RogueCommonPendingAction{Num: 0},
		RogueAeon: &proto.RogueAeon{
			CGAFFPHCNEA: true,
			AeonId:      req.BuffAeonId, // 解锁的命途
		},
		RogueAvatarInfo: &proto.RogueAvatarInfo{
			BaseAvatarIdList: req.BaseAvatarIdList,
			AJJJNLPCEED: &proto.CLPDAOOAHOE{
				MGEFFLOEPBK: &proto.ItemCostList{
					ItemList: []*proto.ItemCost{
						{
							PileItem: &proto.PileItem{
								ItemNum: 80,
								ItemId:  31,
							},
						},
					},
				},
			},
		},
		RoomMap: g.GetRogueMap(),
		RogueVirtualItem: &proto.RogueVirtualItem{
			Money: g.GetDbRogue().CurRogue.CosmicFragment,
		},
		Status: proto.RogueStatus_ROGUE_STATUS_DOING,
	}

	g.GetBattleState().BattleType = spb.BattleType_Battle_ROGUE
	g.Send(cmd.StartRogueScRsp, rsp)
}

// 获取当前模拟宇宙信息
func (g *GamePlayer) GetRogueInfo() *proto.RogueInfo {
	cfRogueManager := gdconf.GetRogueManager()
	rogueInfo := &proto.RogueInfo{
		RogueInfoData: &proto.RogueInfoData{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: cfRogueManager.BeginTime,
				EndTime:   cfRogueManager.EndTime,
				SeasonId:  cfRogueManager.RogueSeason,
			},
			RogueScoreInfo: &proto.RogueScoreRewardInfo{
				BeginTime:            cfRogueManager.BeginTime,
				EndTime:              cfRogueManager.EndTime,
				HasTakenInitialScore: true,
				PoolId:               22,
				PoolRefreshed:        true,
				HasTakenReward:       nil,  // 已领取奖励
				Score:                5000, // 本期得分
			},
			RogueAreaInfo: &proto.RogueAreaInfo{RogueArea: g.GetRogueArea()},
			RogueAeonInfo: &proto.RogueAeonInfo{
				// TODO
				AeonIdList:    []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9},
				IsUnlocked:    true,
				UnlockAeonNum: 9,
			},
		},
	}

	return rogueInfo
}

// 获取模拟宇宙关卡信息
func (g *GamePlayer) GetRogueArea() []*proto.RogueArea {
	rogueAreaList := make([]*proto.RogueArea, 0)
	cfRogueManager := gdconf.GetRogueManager()
	for _, rogueArea := range cfRogueManager.RogueAreaIDList {
		dbRogueArea := g.GetDbRogueArea(rogueArea)
		RogueArea := &proto.RogueArea{
			AreaId:          dbRogueArea.AreaId,
			RogueAreaStatus: proto.RogueAreaStatus(dbRogueArea.RogueAreaStatus),
		}
		rogueAreaList = append(rogueAreaList, RogueArea)
	}

	return rogueAreaList
}

// 新建模拟宇宙
func (g *GamePlayer) NewRogue(avatarIdList []uint32, areaId uint32) {
	// 更新队伍
	if avatarIdList != nil {
		g.GetLineUpById(9).AvatarIdList = avatarIdList
	}
	g.GetLineUp().MainAvatarId = 0

	// 获取地图
	rogueAreaConfig := gdconf.GetRogueAreaConfigById(strconv.Itoa(int(areaId)))
	if rogueAreaConfig == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	rogueMapID := (rogueAreaConfig.AreaProgress * 100) + rogueAreaConfig.Difficulty
	rogueMap := gdconf.GetRogueMapById(rogueMapID)
	if rogueMap == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}

	rogueSceneMap := make(map[uint32]*spb.RogueRoom)
	for id, rogue := range rogueMap.SiteList {
		rogueSceneMap[id] = &spb.RogueRoom{
			RoomId:         gdconf.GetRogueRoomIDBySiteID(id),
			RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
			NextSiteIdList: rogue.NextSiteIDList,
		}
	}

	rogueSceneMap[rogueMap.StartId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY

	g.GetDbRogue().CurRogue = &spb.CurRogue{
		CurAreaId:      areaId,
		CurSiteId:      rogueMap.StartId,
		CosmicFragment: 50,
		RogueSceneMap:  rogueSceneMap,
		RogueMapID:     rogueMapID,
	}
}

func (g *GamePlayer) GetRogueMap() *proto.RogueMapInfo {
	rogue := g.GetDbRogue()
	roomMap := &proto.RogueMapInfo{
		MapId:     rogue.CurRogue.RogueMapID,
		AreaId:    rogue.CurRogue.CurAreaId,
		CurSiteId: rogue.CurRogue.CurSiteId, // 当前id
		CurRoomId: rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomId,
		RoomList:  make([]*proto.RogueRoom, 0),
	}

	for id, rogueScene := range rogue.CurRogue.RogueSceneMap {
		roomList := &proto.RogueRoom{
			SiteId:     id,
			RoomId:     rogueScene.RoomId,
			RoomStatus: proto.RogueRoomStatus(rogueScene.RoomStatus),
		}

		roomMap.RoomList = append(roomMap.RoomList, roomList)
	}

	return roomMap
}

func (g *GamePlayer) GetRogueScene(roomId uint32) (*proto.SceneInfo, map[uint32]*AvatarEntity, map[uint32]*MonsterEntity) {
	rogueRoom := gdconf.GetRogueRoomById(roomId)
	if rogueRoom == nil {
		return nil, nil, nil
	}
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(rogueRoom.MapEntrance)))
	if mapEntrance == nil {
		return nil, nil, nil
	}

	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	scene := &proto.SceneInfo{
		ClientPosVersion:   5,
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

	// 获取场景实体
	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	// 添加角色信息
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	startGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, rogueRoom.GroupID)
	anchor := startGroup.AnchorList[0]
	baseAvatarIdList := g.GetLineUpById(9)
	for id, avatarId := range baseAvatarIdList.AvatarIdList {
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
			avatarEntity[leaderEntityId] = &AvatarEntity{
				AvatarId: avatarId,
				GroupId:  rogueRoom.GroupID,
			}
		} else {
			entityList.EntityId = entityId
			avatarEntity[entityId] = &AvatarEntity{
				AvatarId: avatarId,
				GroupId:  rogueRoom.GroupID,
			}
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)
	for groupID, ida := range rogueRoom.GroupWithContent {
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
		entityGroupLists, x := g.GetRogueNPCMonsterByID(entityGroupLists, sceneGroup, stou32(groupID), monsterEntity, ida)
		monsterEntity = x
		// 添加NPC实体
		entityGroupLists = g.GetNPCByID(entityGroupLists, sceneGroup, stou32(groupID))
		if len(entityGroupLists.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
		}
	}

	return scene, avatarEntity, monsterEntity
}

func (g *GamePlayer) GetRogueNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup, groupID uint32, entityMap map[uint32]*MonsterEntity, ida uint32) (*proto.SceneEntityGroupInfo, map[uint32]*MonsterEntity) {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := uint32(g.GetNextGameObjectGuid())
		rogueMonsterID := gdconf.GetRogueMonsterGroupByGroupID(ida)
		rogueMonster := gdconf.GetRogueMonsterByRogueMonsterID(rogueMonsterID)
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(monsterList.PosX * 1000),
					Y: int32(monsterList.PosY * 1000),
					Z: int32(monsterList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(monsterList.RotY * 1000),
					Z: 0,
				},
			},
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.PlayerPb.WorldLevel,
				MonsterId:  rogueMonster.NpcMonsterID,
				EventId:    rogueMonster.EventID,
			},
		}
		// 添加实体
		entityMap[entityId] = &MonsterEntity{
			MonsterEId: rogueMonster.EventID,
			GroupId:    groupID,
			Pos: &Vector{
				X: int32(monsterList.PosX * 1000),
				Y: int32(monsterList.PosY * 1000),
				Z: int32(monsterList.PosZ * 1000),
			},
			Rot: &Vector{
				X: 0,
				Y: int32(monsterList.RotY * 1000),
				Z: 0,
			},
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList, entityMap
}

// 模拟宇宙攻击事件
func (g *GamePlayer) RogueSceneCastSkillCsReq(rsp *proto.SceneCastSkillScRsp) {
	rsp.BattleInfo.BattleTargetInfo = make(map[uint32]*proto.BattleTargetList)
	rsp.BattleInfo.BattleTargetInfo[2] = &proto.BattleTargetList{
		BattleTargetList: make([]*proto.BattleTarget, 0),
	}
	battleTargetList := make([]*proto.BattleTarget, 0)
	bTL := []uint32{50008, 50007, 50006, 50005, 50004, 50003, 50001, 30072, 30071, 30070, 30069, 30068, 30067, 30066, 30065, 30064, 30063, 30062, 30061, 30060, 30059, 30058, 30057, 30056, 30055, 30054, 30053, 30046, 30045, 30044, 30043, 30042, 30041, 30040, 30039, 30037, 30034, 30024, 30021, 30016, 30012, 30008, 30006, 30004, 30002, 20074, 20069, 20068, 20067, 20066, 20064, 20060, 20058, 20052, 20050, 20049, 20047, 20046, 20045, 20044, 20043, 20040, 20039, 20034, 20033, 20032, 20031, 20024, 20021, 20019, 20018, 20017, 20016, 20015, 20013, 20010, 20009, 20008, 20002}
	for _, id := range bTL {
		battleTarget := &proto.BattleTarget{
			Id: id,
		}
		battleTargetList = append(battleTargetList, battleTarget)
	}
	rsp.BattleInfo.BattleTargetInfo[2].BattleTargetList = battleTargetList
	// 添加角色
	rsp.BattleInfo.BattleAvatarList = g.GetBattleAvatarList(9)
	// 添加buff
	rsp.BattleInfo.BuffList = make([]*proto.BattleBuff, 0)
	for id := range g.GetRogueBuff() {
		battleBuff := &proto.BattleBuff{
			Id:       id,
			Level:    1,
			OwnerId:  4294967295,
			WaveFlag: 4294967295,
		}
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, battleBuff)
	}
	// 路途buff！！！
	rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, &proto.BattleBuff{
		Id:       1000111,
		Level:    1,
		WaveFlag: 4294967295,
		DynamicValues: map[string]float32{
			"SkillIndex": 1,
		},
		TargetIndexList: []uint32{1},
	})
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

// 模拟宇宙攻击事件结算
func (g *GamePlayer) RoguePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq, rsp *proto.PVEBattleResultScRsp) {
	battle := g.GetRogueBattle()[req.BattleId]
	// 队伍状态通知
	g.ChallengeSyncLineupNotify(9)
	rsp.BattleAvatarList = g.GetBattleAvatarList(9)
	// buff同步
	var buffList []uint32
	g.SyncEntityBuffChangeListScNotify(buffList)
	// 积分同步
	g.SyncRogueVirtualItemInfoScNotify()
	// 物品增加通知
	var pileItem []*Material
	pileItem = append(pileItem, &Material{
		Tid: 31,
		Num: 20,
	})
	g.ScenePlaneEventScNotify(pileItem)
	// 祝福选择页通知 SyncRogueCommonPendingActionScNotify
	buffIdList := gdconf.GetBuffListByNum(3)
	g.SyncRogueCommonPendingActionScNotify(buffIdList)
	g.SyncRogueCommonPendingActionScNotify(buffIdList)
	// 场景实体刷新通知 SceneGroupRefreshScNotify （门和删除刚刚战斗的实体）
	// 删除实体
	nitify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
	}
	for _, eventId := range battle.monsterEntityMap {
		entity := g.GetSceneEntity().MonsterEntity[eventId]
		if entity != nil {
			groupRefreshInfo := &proto.SceneGroupRefreshInfo{
				GroupId: entity.GroupId,
				RefreshEntity: []*proto.SceneEntityRefreshInfo{
					{
						DelEntity: eventId,
					},
				},
			}
			nitify.GroupRefreshInfo = append(nitify.GroupRefreshInfo, groupRefreshInfo)
			delete(g.GetSceneEntity().MonsterEntity, eventId)
		}
	}
	// 刷新门
	g.Send(cmd.SceneGroupRefreshScNotify, nitify)

	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

func (g *GamePlayer) QuitRogueCsReq(payloadMsg []byte) {

	g.Send(cmd.QuitRogueScRsp, nil)
}

func (g *GamePlayer) LeaveRogueCsReq(payloadMsg []byte) {
	rsp := &proto.LeaveRogueScRsp{
		RogueInfo: g.GetRogueInfo(),
		Lineup:    g.GetLineUpPb(g.GetLineUp().MainLineUp),
		Scene:     g.GetSceneInfo(g.GetScene().EntryId, g.GetPos(), g.GetRot()),
	}

	g.Send(cmd.LeaveRogueScRsp, rsp)
}

func (g *GamePlayer) HandleRogueCommonPendingActionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.HandleRogueCommonPendingActionCsReq, payloadMsg)
	req := msg.(*proto.HandleRogueCommonPendingActionCsReq)
	if req.BuffSelectResult.BuffId == 0 {
		return
	}
	var buffIdList []uint32
	// 祝福通知
	if g.GetRogue().BuffList[req.BuffSelectResult.BuffId] == nil {
		buffIdList = append(buffIdList, req.BuffSelectResult.BuffId)
	} else {

	}

	g.RogueAddBuff(req.BuffSelectResult.BuffId)

	g.SyncEntityBuffChangeListScNotify(buffIdList)
	// 添加后通知启动
	g.SyncRogueCommonActionResultScNotify(req.BuffSelectResult.BuffId)
	// 模拟宇宙图鉴更新通知？ SyncRogueHandbookDataUpdateScNotify
	// 模拟宇宙常见操作结果通知 SyncRogueCommonActionResultScNotify // add buff, buff状态
	rsp := &proto.HandleRogueCommonPendingActionScRsp{
		Times: 3,
	}

	g.Send(cmd.HandleRogueCommonPendingActionScRsp, rsp)
}
