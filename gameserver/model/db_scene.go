package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type EntityAll interface {
	AddEntity(g *PlayerData, groupID uint32)
}

type SceneMap struct {
	LoadedGroup    map[uint32]*GroupInfo // 已加载场景
	NoLoadedGroup  map[uint32]*GroupInfo // 未加载场景
	SummonUnitInfo *SummonUnitInfo       // 领域
}

type SummonUnitInfo struct {
	EntityId       uint32
	AvatarId       uint32
	AttachEntityId uint32
	SummonUnitId   uint32
	Pos            *proto.Vector
	BuffList       []*OnBuffMap
}

type GroupInfo struct {
	EntryId   uint32               // 地图
	PlaneID   uint32               // 地图2
	FloorID   uint32               // 地图3
	GroupID   uint32               // 区域
	EntityMap map[uint32]EntityAll // 场景实体(加载区域才写！！！
}

type Entity struct {
	EntityId uint32 // 实体id
	InstId   uint32 // 配置表中id
	EntryId  uint32 // 地图
	GroupId  uint32 // 区域
	Pos      *proto.Vector
	Rot      *proto.Vector
	BuffList map[uint32]*OnBuffMap // buff
}

type AvatarEntity struct {
	Entity
	AvatarId   uint32 // 角色id
	LineAvatar *spb.LineAvatarList
}

type MonsterEntity struct {
	Entity
	EventID       uint32 // 怪物id
	PurposeType   string // 类型
	FarmElementID uint32 // 虚影Id
}

type NpcEntity struct {
	Entity
	NpcId uint32 // ncp id
}

type PropEntity struct {
	Entity
	PropId              uint32 // 物品id
	TriggerBattleString uint32 // id
}

func NewSceneMap() *SceneMap { // 清空实体列表用的
	db := &SceneMap{
		LoadedGroup:   make(map[uint32]*GroupInfo),
		NoLoadedGroup: make(map[uint32]*GroupInfo),
	}
	return db
}

func (g *PlayerData) UpSceneMap() {
	db := g.GetOnlineData()
	db.SceneMap = NewSceneMap()
}

func (g *PlayerData) GetSceneMap() *SceneMap {
	db := g.GetOnlineData()
	if db.SceneMap == nil {
		db.SceneMap = NewSceneMap()
	}
	return db.SceneMap
}

func (g *PlayerData) GetSummonUnitInfo() *SummonUnitInfo {
	db := g.GetSceneMap()
	if db.SummonUnitInfo == nil {
		db.SummonUnitInfo = new(SummonUnitInfo)
	}
	return db.SummonUnitInfo
}

func (g *PlayerData) DelSummonUnitInfo() {
	db := g.GetSceneMap()
	db.SummonUnitInfo = new(SummonUnitInfo)
}

func (g *PlayerData) GetLoadedGroup() map[uint32]*GroupInfo {
	db := g.GetSceneMap()
	if db.LoadedGroup == nil {
		db.LoadedGroup = make(map[uint32]*GroupInfo)
	}
	return db.LoadedGroup
}

func (g *PlayerData) GetNoLoadedGroup() map[uint32]*GroupInfo {
	db := g.GetSceneMap()
	if db.NoLoadedGroup == nil {
		db.NoLoadedGroup = make(map[uint32]*GroupInfo)
	}
	return db.NoLoadedGroup
}

func (g *PlayerData) AddLoadedGroup(entryId, planeID, floorID, groupID uint32) {
	db := g.GetLoadedGroup()
	db[groupID] = &GroupInfo{
		EntryId:   entryId,
		PlaneID:   planeID,
		FloorID:   floorID,
		GroupID:   groupID,
		EntityMap: make(map[uint32]EntityAll),
	}
}

func (g *PlayerData) AddNoLoadedGroup(entryId, planeID, floorID, groupID uint32) {
	db := g.GetNoLoadedGroup()
	db[groupID] = &GroupInfo{
		EntryId:   entryId,
		PlaneID:   planeID,
		FloorID:   floorID,
		GroupID:   groupID,
		EntityMap: make(map[uint32]EntityAll),
	}
}

func (g *PlayerData) GetGroupInfoByGroupID(groupID uint32) *GroupInfo {
	db := g.GetLoadedGroup()
	if db[groupID] == nil {
		db[groupID] = &GroupInfo{}
	}
	return db[groupID]
}

func (g *PlayerData) GetEntity(groupID uint32) map[uint32]EntityAll {
	db := g.GetGroupInfoByGroupID(groupID)
	if db.EntityMap == nil {
		db.EntityMap = make(map[uint32]EntityAll)
	}
	return db.EntityMap
}

func (ae *AvatarEntity) AddEntity(g *PlayerData, groupID uint32) {
	db := g.GetEntity(groupID)
	db[ae.EntityId] = ae
}

func (me *MonsterEntity) AddEntity(g *PlayerData, groupID uint32) {
	db := g.GetEntity(groupID)
	db[me.EntityId] = me
}

func (ne *NpcEntity) AddEntity(g *PlayerData, groupID uint32) {
	db := g.GetEntity(groupID)
	db[ne.EntityId] = ne
}

func (pe *PropEntity) AddEntity(g *PlayerData, groupID uint32) {
	db := g.GetEntity(groupID)
	db[pe.EntityId] = pe
}

func (g *PlayerData) AddEntity(groupID uint32, t EntityAll) {
	t.AddEntity(g, groupID)
}

func (g *PlayerData) GetEntityById(id uint32) EntityAll { // 根据实体id拉取实体
	db := g.GetLoadedGroup()
	for _, info := range db {
		if info.EntityMap != nil {
			for eid := range info.EntityMap {
				if eid == id {
					return info.EntityMap[eid]
				}
			}
		}
	}
	return nil
}

func (g *PlayerData) GetEntryId(t EntityAll) uint32 {
	if t == nil {
		return 0
	}
	switch t.(type) {
	case *PropEntity:
		return t.(*PropEntity).EntityId
	case *MonsterEntity:
		return t.(*MonsterEntity).EntityId
	case *NpcEntity:
		return t.(*NpcEntity).EntityId
	default:
		return 0
	}
}

func (g *PlayerData) GetMonsterEntityById(id uint32) *MonsterEntity {
	db := g.GetEntityById(id)
	if db == nil {
		return nil
	}
	switch db.(type) {
	case *MonsterEntity:
		return db.(*MonsterEntity)
	}
	return nil
}

func (g *PlayerData) GetPropEntityById(id uint32) *PropEntity {
	db := g.GetEntityById(id)
	if db == nil {
		return nil
	}
	switch db.(type) {
	case *PropEntity:
		return db.(*PropEntity)
	}
	return nil
}

func (g *PlayerData) GetPropEntity(groupId, instId uint32) *PropEntity {
	db := g.GetEntity(groupId)
	for _, entity := range db {
		switch entity.(type) {
		case *PropEntity:
			if entity.(*PropEntity).GroupId == groupId && entity.(*PropEntity).InstId == instId {
				return entity.(*PropEntity)
			}
		}
	}
	return nil
}

func (g *PlayerData) GetAvatarEntity(id uint32) *AvatarEntity {
	db := g.GetEntityById(id)
	if db == nil {
		return nil
	}
	switch db.(type) {
	case *AvatarEntity:
		return db.(*AvatarEntity)
	}
	return nil
}

func (g *PlayerData) GetCurAvatarEntity() uint32 {
	avatarId := g.GetSceneAvatarId()
	for _, info := range g.GetLoadedGroup() {
		if info.EntityMap != nil {
			for _, e := range info.EntityMap {
				switch x := e.(type) {
				case *AvatarEntity:
					if x.AvatarId == avatarId {
						return x.EntityId
					}
				}
			}
		}
	}
	return 0
}

func (g *PlayerData) GetAllPropEntity() []*PropEntity {
	peList := make([]*PropEntity, 0)
	db := g.GetLoadedGroup()
	for _, groupInfo := range db {
		if groupInfo.EntityMap != nil {
			for _, info := range groupInfo.EntityMap {
				switch info.(type) {
				case *PropEntity:
					peList = append(peList, info.(*PropEntity))
				}
			}
		}
	}
	return peList
}

func (g *PlayerData) GetTriggerBattleString(eventId uint32) *PropEntity {
	db := g.GetLoadedGroup()
	for _, group := range db {
		for _, entity := range group.EntityMap {
			switch entity.(type) {
			case *PropEntity:
				if entity.(*PropEntity).TriggerBattleString == eventId {
					return entity.(*PropEntity)
				}
			}
		}
	}
	return nil
}

func NewScene() *spb.Scene {
	return &spb.Scene{
		EntryId: 2000101,
		Pos:     NewPos(),
		Rot:     NewRot(),
	}
}

func (g *PlayerData) GetScene() *spb.Scene {
	db := g.BasicBin
	if db.Scene == nil {
		db.Scene = NewScene()
	}
	return db.Scene
}

func (g *PlayerData) GetCurEntryId() uint32 {
	db := g.GetScene()
	return db.EntryId
}

func (g *PlayerData) SetCurEntryId(id uint32) {
	db := g.GetScene()
	db.EntryId = id
}

func NewPos() *spb.VectorBin {
	return &spb.VectorBin{
		X: 99,
		Y: 62,
		Z: -4800,
	}
}

func NewRot() *spb.VectorBin {
	return &spb.VectorBin{
		X: 0,
		Y: 0,
		Z: 0,
	}
}

func (g *PlayerData) GetPos() *spb.VectorBin {
	db := g.GetScene()
	if db.Pos == nil {
		db.Pos = NewPos()
	}
	return db.Pos
}

func (g *PlayerData) GetRot() *spb.VectorBin {
	db := g.GetScene()
	if db.Rot == nil {
		db.Rot = NewRot()
	}
	return db.Rot
}

func (g *PlayerData) SetPos(x, y, z int32) {
	db := g.GetPos()
	db.X = x
	db.Y = y
	db.Z = z
}

func (g *PlayerData) SetRot(x, y, z int32) {
	db := g.GetRot()
	db.X = x
	db.Y = y
	db.Z = z
}

func (g *PlayerData) GetPet() *spb.Pet {
	db := g.GetScene()
	if db.Pet == nil {
		db.Pet = new(spb.Pet)
	}
	return db.Pet
}

func (g *PlayerData) SetCurPet(petId uint32) bool {
	db := g.GetPet()
	if petId == 0 {
		db.CurPetId = 0
		return true
	}
	if _, ok := db.UnlockedPetList[petId]; ok {
		db.CurPetId = petId
		return true
	}
	return false
}

func (g *PlayerData) GetCurPet() uint32 {
	return g.GetPet().CurPetId
}

func (g *PlayerData) AddPet(petId uint32) {
	db := g.GetPet()
	if db.UnlockedPetList == nil {
		db.UnlockedPetList = make(map[uint32]bool)
	}
	db.UnlockedPetList[petId] = true
}

func (g *PlayerData) IfLoadMap(levelGroup *gdconf.GoppLevelGroup) bool {
	if levelGroup.GroupName == "TrainVisitorDemo" ||
		levelGroup.GroupName == "TrainVisiter" {
		return false
	}
	switch levelGroup.Category {
	case "": // 基础
		return g.IfMissionLoadMap(levelGroup, true)
	case "Mission": // 任务
		return g.IfMissionLoadMap(levelGroup, true)
	case "BattleProps":
		return g.IfMissionLoadMap(levelGroup, true)
	case "BattleAudiences":
		return g.IfMissionLoadMap(levelGroup, true)
	case "Custom": // 特殊环境场景:模拟宇宙等
		return g.IfMissionLoadMap(levelGroup, false)
	case "System": // 副本/关卡/等入口
		return g.IfMissionLoadMap(levelGroup, true)
	case "Atmosphere": // 特殊交互物品
		return g.IfMissionLoadMap(levelGroup, true)
	default:
		logger.Warn("未知的地图类型 Category:%s", levelGroup.Category)
		return false
	}
}

func (g *PlayerData) IfMissionLoadMap(levelGroup *gdconf.GoppLevelGroup, mainIsLoaded bool) bool {
	c := &LevelGroupMissionConditionSet{
		PlayerData:                    g,
		finishSubMainMissionList:      g.GetFinishSubMainMissionList(),
		subMainMissionList:            g.GetSubMainMissionList(),
		mainMissionList:               g.GetMainMissionList(),
		finishMainMissionList:         g.GetFinishMainMissionList(),
		systemUnlockCondition:         levelGroup.SystemUnlockCondition,
		savedValueCondition:           levelGroup.SavedValueCondition,
		levelGroupMissionConditionSet: nil,
	}

	switch levelGroup.GroupId {
	// case 127, 128, 129:
	// 	return false
	}

	// 检查系统功能解锁条件
	if !c.CheckSystemUnlockCondition() {
		return false
	}

	// FSV 筛选

	if levelGroup.OwnerMainMissionID != 0 {
		if c.mainMissionList[levelGroup.OwnerMainMissionID] == nil {
			return false
		} else {
			mainIsLoaded = true
		}
	}

	missionId := alg.ExtractDigits(levelGroup.GroupName)
	if gdconf.GetSubMainMissionById(missionId) != nil {
		if c.subMainMissionList[missionId] != nil {
			return true
		}
		return false
	}
	if gdconf.GetMainMissionById(missionId) != nil {
		if c.mainMissionList[missionId] != nil {
			return true
		}
		return false
	}

	// 检查加载条件
	if levelGroup.LoadCondition != nil {
		c.levelGroupMissionConditionSet = levelGroup.LoadCondition
		if !c.CheckLevelGroupMissionConditionSet(true) {
			return false
		} else {
			mainIsLoaded = true
		}
	}

	// 检查卸载条件
	if levelGroup.UnloadCondition != nil {
		c.levelGroupMissionConditionSet = levelGroup.UnloadCondition
		if c.CheckLevelGroupMissionConditionSet(false) {
			return false
		}
	}

	// 检查强制卸载条件
	if levelGroup.ForceUnloadCondition != nil {
		c.levelGroupMissionConditionSet = levelGroup.ForceUnloadCondition
		if c.CheckLevelGroupMissionConditionSet(false) {
			return false
		}
	}

	return mainIsLoaded
}

type LevelGroupMissionConditionSet struct {
	*PlayerData
	finishSubMainMissionList      map[uint32]*spb.MissionInfo // 已完成子任务
	subMainMissionList            map[uint32]*spb.MissionInfo // 接受的子任务
	mainMissionList               map[uint32]*spb.MissionInfo // 接取的主任务
	finishMainMissionList         map[uint32]*spb.MissionInfo // 已完成的主任务
	systemUnlockCondition         *gdconf.LevelGroupSystemUnlockConditionSet
	savedValueCondition           *gdconf.LevelGroupSavedValueConditionSet
	levelGroupMissionConditionSet *gdconf.LevelGroupMissionConditionSet
}

func (c *LevelGroupMissionConditionSet) CheckSystemUnlockCondition() bool {
	if c.systemUnlockCondition == nil {
		return true
	}
	result := c.systemUnlockCondition.Operation != constant.LogicOperationTypeOr
	for _, conditionId := range c.systemUnlockCondition.Conditions {
		gsuk := gdconf.GetGroupSystemUnlockData(conditionId)
		if gsuk == nil {
			continue
		}
		part := c.GetUnlockStatus(gdconf.GetFuncUnlockDataConditions(gsuk.UnlockID))
		if c.systemUnlockCondition.Operation == constant.LogicOperationTypeOr && part {
			return true
		}
		if c.systemUnlockCondition.Operation == constant.LogicOperationTypeAnd && !part {
			return false
		}

		if c.systemUnlockCondition.Operation != constant.LogicOperationTypeNot || !part {
			continue
		}
		return false
	}
	return result
}

func (g *PlayerData) GetUnlockStatus(conditions []*gdconf.ConditionParam) bool {
	if conditions == nil {
		return false
	}
	for _, condition := range conditions {
		switch condition.Type {
		case constant.ConditionTypeFinishMainMission:
			if !g.GetIsJumpMission() &&
				g.GetFinishMainMissionById(alg.S2U32(condition.Param)) == nil {
				return false
			}
		case constant.ConditionTypePlayerLevel:
			if g.GetLevel() < alg.S2U32(condition.Param) {
				return false
			}
		case constant.ConditionTypeWorldLevel:
			if g.GetWorldLevel() < alg.S2U32(condition.Param) {
				return false
			}
		case constant.ConditionTypeFinishSubMission:
			if !g.GetIsJumpMission() &&
				g.GetFinishSubMainMissionById(alg.S2U32(condition.Param)) == nil {
				return false
			}
		case constant.ConditionTypeInStoryLine:
			storyLine := g.GetCurChangeStoryInfo()
			if storyLine != nil &&
				storyLine.ChangeStoryId != alg.S2U32(condition.Param) {
				return false
			}
		}
	}
	return true
}

func (c *LevelGroupMissionConditionSet) CheckLevelGroupMissionConditionSet(defaultResult bool) bool {
	if len(c.levelGroupMissionConditionSet.Conditions) == 0 {
		return defaultResult
	}
	isLoaded := false // c.levelGroupMissionConditionSet.Operation == constant.LogicOperationTypeAnd
	for _, condition := range c.levelGroupMissionConditionSet.Conditions {
		status := c.GetMissionStatus(condition.ID)
		if c.GetIsJumpMission() {
			status = constant.LevelGroupMissionPhaseFinish
		}
		newPhase := func() constant.LevelGroupMissionPhase {
			if condition.LevelGroupMissionPhase == constant.LevelGroupMissionPhaseCancel {
				return constant.LevelGroupMissionPhaseFinish
			}
			return condition.LevelGroupMissionPhase
		}()
		if status != newPhase {
			isLoaded = false
			if c.levelGroupMissionConditionSet.Operation ==
				constant.LogicOperationTypeAnd {
				break
			}
		} else {
			isLoaded = true
			if c.levelGroupMissionConditionSet.Operation ==
				constant.LogicOperationTypeOr {
				break
			}
		}
	}
	return isLoaded
}

func (g *PlayerData) GetMissionStatus(id uint32) constant.LevelGroupMissionPhase {
	if g.GetSubMainMissionById(id) != nil ||
		g.GetMainMissionById(id) != nil {
		return constant.LevelGroupMissionPhaseAccept
	}
	if g.GetFinishSubMainMissionById(id) != nil ||
		g.GetFinishMainMissionById(id) != nil {
		return constant.LevelGroupMissionPhaseFinish
	}
	return constant.LevelGroupMissionPhaseCancel
}

// 检查场景上是否有实体需要卸载/加载
func (g *PlayerData) AutoEntryGroup() ([]*GroupInfo, []*GroupInfo) {
	loadedGroup := g.GetLoadedGroup()        // 已加载区域
	noLoadedGroup := g.GetNoLoadedGroup()    // 未加载区域
	uninstallGroup := make([]*GroupInfo, 0)  // 卸载场景列表
	loadedGroupList := make([]*GroupInfo, 0) // 加载列表
	// 检查已加载区域是否需要卸载
	for id, info := range loadedGroup {
		group := gdconf.GetServerGroupById(info.PlaneID, info.FloorID, info.GroupID)
		if group == nil {
			continue
		}
		if !g.IfLoadMap(group) {
			// 此处卸载
			uninstallGroup = append(uninstallGroup, info)
			noLoadedGroup[id] = info
			delete(loadedGroup, id)
		}
	}
	// 检查未加载区域是否需要加载
	for id, info := range noLoadedGroup {
		group := gdconf.GetServerGroupById(info.PlaneID, info.FloorID, info.GroupID)
		if group == nil {
			continue
		}
		if g.IfLoadMap(group) {
			// 此处加载
			loadedGroupList = append(loadedGroupList, info)
			loadedGroup[id] = info
			delete(noLoadedGroup, id)
		}
	}

	// 卸载/加载
	return uninstallGroup, loadedGroupList
}

func NewBlockMap() map[uint32]*spb.BlockBin {
	return map[uint32]*spb.BlockBin{}
}

func (g *PlayerData) GetBlockMap() map[uint32]*spb.BlockBin {
	db := g.GetOnlineData()
	if db.BlockMap == nil {
		db.BlockMap = NewBlockMap()
	}
	return db.BlockMap
}

func (g *PlayerData) GetAllBlockMap() map[uint32]*spb.BlockBin {
	db := g.GetOnlineData()
	if db.BlockMap == nil {
		db.BlockMap = NewBlockMap()
	}
	blockMap := make(map[uint32]*spb.BlockBin)
	for k, v := range db.BlockMap {
		blockMap[k] = v
	}
	return blockMap
}

// 从db拉取地图数据
func (g *PlayerData) GetBlock(entryId uint32) *spb.BlockBin {
	newEntryId := entryId
	if entryId >= 10000000 {
		newEntryId = alg.S2U32(strconv.Itoa(int(entryId))[:7])
	}
	if mapEntrance := gdconf.GetMapEntranceById(newEntryId); mapEntrance == nil {
		newEntryId = entryId
	}
	db := g.GetBlockMap()
	if db[newEntryId] == nil {
		bin := database.GetBlockData(database.GSS.PlayerDataMysql,
			g.GetBasicBin().Uid, newEntryId)
		block := new(spb.BlockBin)
		if err := pb.Unmarshal(bin.BinData, block); err != nil {
			logger.Debug("entryId:%v,block error", newEntryId)
		}
		db[newEntryId] = block
	}

	db[newEntryId].EntryId = newEntryId
	return db[newEntryId]
}

// 更新地图数据到数据库
func (g *PlayerData) UpdateBlock(block *spb.BlockBin) {
	bin, err := pb.Marshal(block)
	if err != nil {
		return
	}
	newEntryId := block.EntryId
	if block.EntryId >= 10000000 {
		newEntryId = alg.S2U32(strconv.Itoa(int(block.EntryId))[:7])
	}
	if mapEntrance := gdconf.GetMapEntranceById(newEntryId); mapEntrance == nil {
		newEntryId = block.EntryId
	}
	blockData := &constant.BlockData{
		Uid:         g.GetBasicBin().Uid,
		EntryId:     newEntryId,
		DataVersion: 0, // TODO
		BinData:     bin,
	}
	if err = database.UpdateBlockData(database.GSS.PlayerDataMysql,
		blockData); err != nil {
		logger.Debug("updata block data error:%s", err.Error())
	}
}

func (g *PlayerData) GetPropState(db *spb.BlockBin, groupId, propId uint32, state string) uint32 {
	if db == nil {
		return gdconf.GetStateValue(state)
	}
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	if db.BlockList[groupId].PropInfo == nil {
		db.BlockList[groupId].PropInfo = make(map[uint32]*spb.PropInfo)
	}
	if db.BlockList[groupId].PropInfo[propId] == nil {
		db.BlockList[groupId].PropInfo[propId] = &spb.PropInfo{
			InstId: propId,
			// PropId:    prop.PropID,
			PropState: gdconf.GetStateValue(state),
		}
	}
	return db.BlockList[groupId].PropInfo[propId].PropState
}

func (g *PlayerData) UpPropState(db *spb.BlockBin, groupId, propId, state uint32) {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	if db.BlockList[groupId].PropInfo == nil {
		db.BlockList[groupId].PropInfo = make(map[uint32]*spb.PropInfo)
	}
	if db.BlockList[groupId].PropInfo[propId] == nil {
		db.BlockList[groupId].PropInfo[propId] = &spb.PropInfo{
			InstId:    propId,
			PropState: state,
		}
	} else {
		db.BlockList[groupId].PropInfo[propId].PropState = state
	}
}

func (g *PlayerData) GetGroupState(db *spb.BlockBin, groupId uint32) uint32 {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	return db.BlockList[groupId].GroupState
}

func (g *PlayerData) SetGroupState(db *spb.BlockBin, groupId, groupState uint32) {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	db.BlockList[groupId].GroupState = groupState
}

func (g *PlayerData) GetFloorSavedData(entryId uint32) map[string]int32 {
	db := g.GetBlock(entryId)
	if db.FloorSavedData == nil {
		db.FloorSavedData = make(map[string]int32)
	}
	planeID, floorID, ok := gdconf.GetPFlaneID(entryId)
	if !ok {
		logger.Debug(text.GetTextByL(g.GetLanguageType(), 76), entryId)
		return nil
	}
	dimensionInfo := gdconf.GetGoppRtLevelDimensionInfo(planeID, floorID, g.GetDimensionId())
	if dimensionInfo == nil {
		return db.FloorSavedData
	}
	for _, savedValue := range dimensionInfo.SavedValues {
		if savedValue.Name == "Connection_MuteStairs" {
			db.FloorSavedData[savedValue.Name] = 0
		} else if strings.Contains(savedValue.Name, "Activity") {
			db.FloorSavedData[savedValue.Name] = savedValue.MaxValue
		} else if strings.Contains(savedValue.Name, "Build") {
			db.FloorSavedData[savedValue.Name] = 1
		} else if strings.Contains(savedValue.Name, "Progress") {
			db.FloorSavedData[savedValue.Name] = 100
		} else if strings.Contains(savedValue.Name, "Onboarded") {
			db.FloorSavedData[savedValue.Name] = 1
		}
		if _, ok = db.FloorSavedData[savedValue.Name]; !ok {
			db.FloorSavedData[savedValue.Name] = savedValue.DefaultValue
		}
	}
	return db.FloorSavedData
}

func (g *PlayerData) SetFloorSavedData(entryId uint32, key string, v int32) {
	db := g.GetBlock(entryId)
	if db.FloorSavedData == nil {
		db.FloorSavedData = make(map[string]int32)
	}
	db.FloorSavedData[key] = v
}

func (g *PlayerData) ObjectCaptureUpPropState(db *spb.BlockBin, groupId, propId, state uint32) {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	if db.BlockList[groupId].PropInfo == nil {
		db.BlockList[groupId].PropInfo = make(map[uint32]*spb.PropInfo)
	}
	if db.BlockList[groupId].PropInfo[propId] == nil {
		db.BlockList[groupId].PropInfo[propId] = &spb.PropInfo{
			InstId:    propId,
			PropState: state,
		}
	}
}

func (g *PlayerData) StageObjectCapture(sceneGroup *gdconf.GoppLevelGroup, prop *gdconf.PropList, groupId uint32, db *spb.BlockBin) {
	if db == nil {
		return
	}
	if strings.Contains(prop.Name, "Elevator0") {
		g.ObjectCaptureUpPropState(db, groupId, prop.ID, 15)
		return
	}
	if strings.Contains(prop.Name, "Door") {
		g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
		return
	}
	if prop.ValueSource != nil {
		for _, v := range prop.ValueSource.Values {
			if v.Key == "IsAutoDoor" {
				g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
				break
			}
			if strings.Contains(v.Key, "ElevatorLock") {
				g.ObjectCaptureUpPropState(db, groupId, prop.ID, 15)
				break
			}
		}
	}
	if prop.StageObjectCapture != nil { // 特殊处理
		switch prop.StageObjectCapture.BlockAlias {
		case "RogueLobby_01": // 模拟宇宙入口直接开放
			g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
			break
		}
	}
	if sceneGroup.GroupName == "Rouge" && gdconf.GetStateValue(prop.State) == 0 {
		g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
		return
	}
	if conf := gdconf.GetSpecialProp(db.EntryId); conf != nil {
		if spGroup := conf.GroupList[groupId]; spGroup != nil {
			if state := spGroup.PropState[prop.ID]; state != "" {
				g.ObjectCaptureUpPropState(db, groupId, prop.ID, gdconf.GetStateValue(state))
			}
		}
	}
}

func FloorTentry(floorID uint32) uint32 {
	if floorID < 10000000 {
		return 1000001
	}
	st := strconv.Itoa(int(floorID))
	entryId := alg.S2U32(st[:6] + st[7:8])
	if mapEntrance := gdconf.GetMapEntranceById(entryId); mapEntrance == nil {
		return floorID
	}
	return entryId
}

/****************************************************功能***************************************************/

func (g *PlayerData) GetPosPb() *proto.Vector {
	db := g.GetPos()
	return &proto.Vector{
		Y: db.Y,
		X: db.X,
		Z: db.Z,
	}
}

func (g *PlayerData) GetRotPb() *proto.Vector {
	db := g.GetRot()
	return &proto.Vector{
		Y: db.Y,
		X: db.X,
		Z: db.Z,
	}
}

func (g *PlayerData) GetPropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, db *spb.BlockBin, entryId uint32) *proto.SceneEntityGroupInfo {
	for _, propList := range sceneGroup.PropList {
		entityId := g.GetNextGameObjectGuid()
		g.StageObjectCapture(sceneGroup, propList, sceneGroup.GroupId, db)
		propState := g.GetPropState(db, sceneGroup.GroupId, propList.ID, propList.State)
		pos := &proto.Vector{
			X: int32(propList.PosX * 1000),
			Y: int32(propList.PosY * 1000),
			Z: int32(propList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(propList.RotX * 1000),
			Y: int32(propList.RotY * 1000),
			Z: int32(propList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId, // 文件名后那个G
			InstId:   propList.ID,        // ID
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			EntityOneofCase: &proto.SceneEntityInfo_Prop{
				Prop: &proto.ScenePropInfo{
					PropId:    propList.PropID, // PropID
					PropState: propState,
				},
			},
		}
		// 添加物品实体
		g.AddEntity(sceneGroup.GroupId, &PropEntity{
			Entity: Entity{
				EntityId: entityId,
				InstId:   propList.ID,
				EntryId:  entryId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			PropId:              propList.PropID,
			TriggerBattleString: propList.TriggerBattleString,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *PlayerData) GetNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) *proto.SceneEntityGroupInfo {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(monsterList.PosX * 1000),
			Y: int32(monsterList.PosY * 1000),
			Z: int32(monsterList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(monsterList.RotX * 1000),
			Y: int32(monsterList.RotY * 1000),
			Z: int32(monsterList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
				NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.GetWorldLevel(),
					MonsterId:  monsterList.NPCMonsterID,
					EventId:    monsterList.EventID,
				},
			},
		}
		// 添加怪物实体
		g.AddEntity(sceneGroup.GroupId, &MonsterEntity{
			Entity: Entity{
				InstId:   monsterList.ID,
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			EventID:       monsterList.EventID,
			PurposeType:   monsterList.PurposeType,
			FarmElementID: monsterList.FarmElementID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *PlayerData) GetSceneInfo(entryId uint32, pos, rot *proto.Vector, lineUp *spb.Line) *proto.SceneInfo {
	planeID, floorID, ok := gdconf.GetPFlaneID(entryId)
	if !ok {
		logger.Debug(text.GetTextByL(g.GetLanguageType(), 76), entryId)
		entryId = 2000101
		planeID, floorID, _ = gdconf.GetPFlaneID(entryId)
	}
	leaderEntityId := g.GetNextGameObjectGuid()
	dimensionId := g.GetDimensionId()
	gameModeType := gdconf.GetPlaneType(planeID)
	scene := &proto.SceneInfo{
		MGLHEBHJABE:        make([]uint32, 0),
		ClientPosVersion:   0,
		WorldId:            gdconf.GetWorldId(planeID),
		LeaderEntityId:     leaderEntityId,
		FloorId:            floorID,
		GameModeType:       gameModeType,
		PlaneId:            planeID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		SceneMissionInfo: &proto.MissionStatusBySceneInfo{
			DisabledMainMissionIdList:   make([]uint32, 0),
			FinishedMainMissionIdList:   make([]uint32, 0),
			SubMissionStatusList:        make([]*proto.Mission, 0),
			UnfinishedMainMissionIdList: make([]uint32, 0),
			MainMissionMcvList:          make([]*proto.MainMissionCustomValue, 0),
		},
		FloorSavedData:     g.GetFloorSavedData(entryId),
		GameStoryLineId:    g.GameStoryLineId(),
		DimensionId:        dimensionId,
		EntityBuffInfoList: make([]*proto.EntityBuffInfo, 0),
	}
	for i := uint32(0); i < 7; i++ {
		scene.LightenSectionList = append(scene.LightenSectionList, i)
	}
	// 获取场景实体
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 清理老实体列表
	g.UpSceneMap()
	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	blockBin := g.GetBlock(entryId)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)
	missionMap := make(map[uint32]bool, 0)
	foorMap := gdconf.GetServerGroup(planeID, floorID)
	if foorMap == nil {
		return nil
	}
	dimensionInfo := gdconf.GetGoppRtLevelDimensionInfo(planeID, floorID, dimensionId)
	for _, levelGroup := range foorMap {
		if !dimensionInfo.GroupIndexList[levelGroup.Index] {
			continue
		}
		// 添加GetMissionStatusBySceneInfo
		addMainMission := func(mainMissionId uint32) {
			if mainMissionId != 0 &&
				!missionMap[mainMissionId] {
				missionMap[mainMissionId] = true
				status := g.GetMissionStatus(mainMissionId)
				if g.GetIsJumpMission() {
					status = constant.LevelGroupMissionPhaseFinish
				}
				if status == constant.LevelGroupMissionPhaseFinish {
					scene.SceneMissionInfo.FinishedMainMissionIdList = append(scene.SceneMissionInfo.FinishedMainMissionIdList,
						mainMissionId)
				} else {
					scene.SceneMissionInfo.UnfinishedMainMissionIdList = append(scene.SceneMissionInfo.UnfinishedMainMissionIdList,
						mainMissionId)
				}
			}
		}
		conditionSet := func(set *gdconf.LevelGroupMissionConditionSet) {
			if set == nil {
				return
			}
			for _, condition := range set.Conditions {
				if !missionMap[condition.ID] {
					continue
				}
				switch condition.LevelGroupMissionType {
				case constant.LevelGroupMissionTypeSubMission:
					scene.SceneMissionInfo.SubMissionStatusList = append(scene.SceneMissionInfo.SubMissionStatusList,
						g.getProtoMission(condition.ID))
				case constant.LevelGroupMissionTypeEventMission:
					addMainMission(levelGroup.OwnerMainMissionID)
				}
				missionMap[condition.ID] = true
			}
		}
		addMainMission(levelGroup.OwnerMainMissionID)
		conditionSet(levelGroup.LoadCondition)
		conditionSet(levelGroup.UnloadCondition)
		conditionSet(levelGroup.ForceUnloadCondition)

		if subMissionId := alg.ExtractDigits(levelGroup.GroupName); subMissionId != 0 &&
			!missionMap[subMissionId] {
			missionMap[subMissionId] = true
			scene.SceneMissionInfo.SubMissionStatusList = append(scene.SceneMissionInfo.SubMissionStatusList,
				g.getProtoMission(subMissionId))
		}

		if !g.IfLoadMap(levelGroup) {
			g.AddNoLoadedGroup(entryId, planeID, floorID, levelGroup.GroupId)
			continue
		} else {
			// logger.Info("加载组PlaneID:%v,FloorID:%v,GroupId:%v,Index:%v",
			// 	planeID, floorID, levelGroup.GroupId, levelGroup.Index)
			g.AddLoadedGroup(entryId, planeID, floorID, levelGroup.GroupId)
		}
		// scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)
		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
			State:      g.GetGroupState(blockBin, levelGroup.GroupId),
		}
		// 添加物品实体
		g.GetPropByID(entityGroupLists, levelGroup, blockBin, entryId)
		// 添加怪物实体
		g.GetNPCMonsterByID(entityGroupLists, levelGroup)
		// 添加NPC实体
		g.GetSceneNPCByConf(entityGroupLists, levelGroup)
		if len(entityGroupLists.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
		}
	}
	return scene
}

func (g *PlayerData) getProtoMission(subMissionId uint32) *proto.Mission {
	if gdconf.GetSubMainMissionById(subMissionId) == nil {
		return nil
	}
	status := g.GetMissionStatus(subMissionId)
	if g.GetIsJumpMission() {
		status = constant.LevelGroupMissionPhaseFinish
	}
	info := &proto.Mission{
		Id: subMissionId,
	}
	switch status {
	case constant.LevelGroupMissionPhaseFinish:
		info.Status = proto.MissionStatus_MISSION_FINISH
	case constant.LevelGroupMissionPhaseAccept:
		info.Status = proto.MissionStatus_MISSION_DOING
	case constant.LevelGroupMissionPhaseCancel:
		info.Status = proto.MissionStatus_MISSION_PREPARED
	}

	return info
}

/********************************************怪物******************************************/

// 添加怪物
func (g *PlayerData) AddMonsterSceneEntityRefreshInfo(mazeGroupID uint32, monsterList map[uint32]*gdconf.MonsterList) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, monster := range monsterList {
		entityId := g.GetNextGameObjectGuid()
		monsterPos := &proto.Vector{
			X: int32(monster.PosX * 1000),
			Y: int32(monster.PosY * 1000),
			Z: int32(monster.PosZ * 1000),
		}
		monsterRot := &proto.Vector{
			X: int32(monster.RotX * 1000),
			Y: int32(monster.RotY * 1000),
			Z: int32(monster.RotZ * 1000),
		}
		seri := &proto.SceneEntityRefreshInfo{
			Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					GroupId:  mazeGroupID,
					InstId:   monster.ID,
					EntityId: entityId,
					Motion: &proto.MotionInfo{
						Pos: monsterPos,
						Rot: monsterRot,
					},
					EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
						NpcMonster: &proto.SceneNpcMonsterInfo{
							WorldLevel: g.GetWorldLevel(),
							MonsterId:  monster.NPCMonsterID,
							EventId:    monster.EventID,
						},
					},
				},
			},
		}
		// 添加怪物实体
		g.AddEntity(mazeGroupID, &MonsterEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  mazeGroupID,
				Pos:      monsterPos,
				Rot:      monsterRot,
				InstId:   monster.ID,
			},
			EventID:       monster.EventID,
			PurposeType:   monster.PurposeType,
			FarmElementID: monster.FarmElementID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 删除怪物
func (g *PlayerData) GetDelSceneGroupRefreshInfo(mem []uint32) []*proto.GroupRefreshInfo {
	sceneGroupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	for _, id := range mem {
		entity := g.GetMonsterEntityById(id)
		if entity == nil {
			continue
		}
		sgri := &proto.GroupRefreshInfo{
			State:   0,
			GroupId: entity.GroupId,
			RefreshEntity: []*proto.SceneEntityRefreshInfo{
				{
					Refresh: &proto.SceneEntityRefreshInfo_DeleteEntity{
						DeleteEntity: entity.EntityId,
					},
				},
			},
			RefreshType: proto.SceneGroupRefreshType_SCENE_GROUP_REFRESH_TYPE_LOADED,
		}
		sceneGroupRefreshInfo = append(sceneGroupRefreshInfo, sgri)
	}
	return sceneGroupRefreshInfo
}

/********************************************NPC******************************************/

func (g *PlayerData) GetSceneNPCByConf(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) *proto.SceneEntityGroupInfo {
	for _, npcList := range sceneGroup.NPCList {
		pos := &proto.Vector{
			X: int32(npcList.PosX * 1000),
			Y: int32(npcList.PosY * 1000),
			Z: int32(npcList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(npcList.RotX * 1000),
			Y: int32(npcList.RotY * 1000),
			Z: int32(npcList.RotZ * 1000),
		}
		// 添加npc
		npcEntity := &NpcEntity{
			Entity: Entity{
				EntityId: g.GetNextGameObjectGuid(),
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
				InstId:   npcList.ID,
			},
			NpcId: npcList.NPCID,
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList,
			g.GetNpcSceneEntityInfo(npcEntity))
	}
	return entityGroupList
}

// 添加Npc
func (g *PlayerData) AddNpcSceneEntityRefreshInfo(mazeGroupID uint32, npcList map[uint32]*gdconf.NPCList) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, npc := range npcList {
		pos := &proto.Vector{
			X: int32(npc.PosX * 1000),
			Y: int32(npc.PosY * 1000),
			Z: int32(npc.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(npc.RotX * 1000),
			Y: int32(npc.RotY * 1000),
			Z: int32(npc.RotZ * 1000),
		}
		// 添加Npc实体
		npcEntity := &NpcEntity{
			Entity: Entity{
				EntityId: g.GetNextGameObjectGuid(),
				GroupId:  mazeGroupID,
				Pos:      pos,
				Rot:      rot,
				InstId:   npc.ID,
			},
			NpcId: npc.NPCID,
		}
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo,
			g.GetNpcSceneEntityRefreshInfo(npcEntity))
	}
	return sceneEntityRefreshInfo
}

func (g *PlayerData) GetNpcSceneEntityInfo(npcEntity *NpcEntity) *proto.SceneEntityInfo {
	if npcEntity == nil {
		return nil
	}
	info := &proto.SceneEntityInfo{
		InstId:   npcEntity.InstId,
		GroupId:  npcEntity.GroupId,
		EntityId: npcEntity.EntityId,
		Motion: &proto.MotionInfo{
			Pos: npcEntity.Pos,
			Rot: npcEntity.Rot,
		},
		EntityOneofCase: &proto.SceneEntityInfo_Npc{
			Npc: &proto.SceneNpcInfo{
				ExtraInfo: nil,
				NpcId:     npcEntity.NpcId,
			},
		},
	}
	g.AddEntity(npcEntity.GroupId, npcEntity)

	return info
}

func (g *PlayerData) GetNpcSceneEntityRefreshInfo(npcEntity *NpcEntity) *proto.SceneEntityRefreshInfo {
	info := &proto.SceneEntityRefreshInfo{
		Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
			AddEntity: g.GetNpcSceneEntityInfo(npcEntity),
		},
	}

	return info
}

/********************************************物品******************************************/

// 添加物品实体
func (g *PlayerData) AddPropSceneEntityRefreshInfo(group *gdconf.GoppLevelGroup, mazeGroupID uint32, propList map[uint32]*gdconf.PropList, db *spb.BlockBin) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, prop := range propList {
		g.StageObjectCapture(group, prop, mazeGroupID, db)
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(prop.PosX * 1000),
			Y: int32(prop.PosY * 1000),
			Z: int32(prop.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(prop.RotX * 1000),
			Y: int32(prop.RotY * 1000),
			Z: int32(prop.RotZ * 1000),
		}
		seri := &proto.SceneEntityRefreshInfo{
			Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					GroupId:  mazeGroupID,
					InstId:   prop.ID,
					EntityId: entityId,
					Motion: &proto.MotionInfo{
						Pos: pos,
						Rot: rot,
					},
					EntityOneofCase: &proto.SceneEntityInfo_Prop{
						Prop: &proto.ScenePropInfo{
							PropId:    prop.PropID, // PropID
							PropState: g.GetPropState(db, mazeGroupID, prop.ID, prop.State),
						},
					},
				},
			},
		}
		// 添加物品实体
		g.AddEntity(mazeGroupID, &PropEntity{
			Entity: Entity{
				EntityId: entityId,
				InstId:   prop.ID,
				EntryId:  db.EntryId,
				GroupId:  mazeGroupID,
				Pos:      pos,
				Rot:      rot,
			},
			PropId:              prop.PropID,
			TriggerBattleString: prop.TriggerBattleString,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

/********************************************角色******************************************/

func (g *PlayerData) GetSceneAvatarByLineUP(entityGroupList *proto.SceneEntityGroupInfo, lineUp *spb.Line, leaderEntityId uint32, pos, rot *proto.Vector) {
	if lineUp.AvatarIdList == nil {
		lineUp.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	}
	if lineUp.AvatarIdList[lineUp.LeaderSlot] == nil {
		lineUp.AvatarIdList[lineUp.LeaderSlot] = &spb.LineAvatarList{
			Slot:           lineUp.LeaderSlot,
			AvatarId:       8001,
			LineAvatarType: spb.AvatarType_AVATAR_FORMAL_TYPE,
		}
	}

	for sole, line := range lineUp.AvatarIdList {
		if line.AvatarId == 0 {
			continue
		}
		entityId := leaderEntityId
		if sole != lineUp.LeaderSlot {
			entityId = g.GetNextGameObjectGuid()
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList,
			g.GetAvatarSceneEntityInfo(line, entityId, pos, rot))
	}
}

// 添加角色
func (g *PlayerData) GetAddAvatarSceneEntityRefreshInfo(lineUp *spb.Line, pos, rot *proto.Vector) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, line := range lineUp.AvatarIdList {
		if line.AvatarId == 0 {
			continue
		}
		entityId := g.GetNextGameObjectGuid()
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo,
			g.GetAvatarSceneEntityRefreshInfo(line, entityId, pos, rot))
	}
	return sceneEntityRefreshInfo
}

func (g *PlayerData) GetSceneGroupRefreshInfoByLineUP(lineUp *spb.Line, pos, rot *proto.Vector) []*proto.GroupRefreshInfo {
	groupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	for _, line := range lineUp.AvatarIdList {
		if line.AvatarId == 0 {
			continue
		}
		entityId := g.GetNextGameObjectGuid()
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity,
			g.GetAvatarSceneEntityRefreshInfo(line, entityId, pos, rot))
	}
	groupRefreshInfo = append(groupRefreshInfo, sceneGroupRefreshInfo)
	return groupRefreshInfo
}

func (g *PlayerData) GetAvatarSceneEntityRefreshInfo(line *spb.LineAvatarList, entityId uint32, pos, rot *proto.Vector) *proto.SceneEntityRefreshInfo {
	info := &proto.SceneEntityRefreshInfo{
		Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
			AddEntity: g.GetAvatarSceneEntityInfo(line, entityId, pos, rot),
		},
	}

	return info
}

func (g *PlayerData) GetAvatarSceneEntityInfo(line *spb.LineAvatarList, entityId uint32, pos, rot *proto.Vector) *proto.SceneEntityInfo {
	info := &proto.SceneEntityInfo{
		EntityOneofCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType(line.LineAvatarType),
			BaseAvatarId: gdconf.SpecialAvatarGetBaseAvatarID(line.AvatarId),
			MapLayer:     0,
			Uid:          line.Uid,
		}},
		Motion: &proto.MotionInfo{
			Pos: pos,
			Rot: rot,
		},
		EntityId: entityId,
	}
	g.AddEntity(0, &AvatarEntity{
		Entity: Entity{
			EntityId: entityId,
			GroupId:  0,
			Pos:      pos,
			Rot:      rot,
		},
		AvatarId:   gdconf.SpecialAvatarGetPlayerID(line.AvatarId),
		LineAvatar: line,
	})

	return info
}

/**************************************领域******************************/

// 添加领域
func (g *PlayerData) GetAddBuffSceneEntityRefreshInfo(casterEntityId, summonId, entityId uint32, pos *proto.Vector) []*proto.GroupRefreshInfo {
	groupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
		Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
			AddEntity: &proto.SceneEntityInfo{
				Motion: &proto.MotionInfo{
					Pos: pos,
					Rot: &proto.Vector{Y: 139439},
				},
				EntityId: entityId,
				EntityOneofCase: &proto.SceneEntityInfo_SummonUnit{
					SummonUnit: &proto.SceneSummonUnitInfo{
						CasterEntityId:  casterEntityId,
						AttachEntityId:  casterEntityId,
						SummonUnitId:    summonId,
						CreateTimeMs:    uint64(time.Now().UnixMilli()),
						TriggerNameList: make([]string, 0),
						LifeTimeMs:      -1,
					},
				},
			},
		},
	}
	sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	groupRefreshInfo = append(groupRefreshInfo, sceneGroupRefreshInfo)
	return groupRefreshInfo
}
