package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
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
	EventID uint32 // 怪物id
}

type NpcEntity struct {
	Entity
	NpcId uint32 // ncp id
}

type PropEntity struct {
	Entity
	PropId uint32 // 物品id
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

func (g *PlayerData) IfLoadMap(levelGroup *gdconf.GoppLevelGroup) bool {
	if levelGroup.GroupName == "TrainVisitorDemo" {
		return false
	}
	switch levelGroup.Category {
	case "": // 基础
		return g.IfMissionLoadMap(levelGroup, true)
	case "System": // 副本/关卡/等入口
		return g.IfMissionLoadMap(levelGroup, true)
	case "Atmosphere": // 特殊交互物品
		return true
	case "Custom": // 特殊环境场景:模拟宇宙等
		return false
	case "Mission": // 任务
		return g.IfMissionLoadMap(levelGroup, false)
	default:
		logger.Warn("未知的地图类型 Category:%s", levelGroup.Category)
		return true
	}
}

func (g *PlayerData) IfMissionLoadMap(levelGroup *gdconf.GoppLevelGroup, mainIsLoaded bool) bool {
	finishSubMainMissionList := g.GetFinishSubMainMissionList() // 已完成子任务
	subMainMissionList := g.GetSubMainMissionList()             // 接受的子任务
	mainMissionList := g.GetMainMissionList()                   // 接取的主任务
	finishMainMissionList := g.GetFinishMainMissionList()       // 已完成的主任务
	isLoaded := mainIsLoaded
	if levelGroup.OwnerMainMissionID != 0 {
		if mainMissionList[levelGroup.OwnerMainMissionID] == nil {
			return false
		}
	}
	if levelGroup.LoadCondition == nil &&
		levelGroup.UnloadCondition == nil {
		if levelGroup.Category == "Mission" && levelGroup.OwnerMainMissionID != 0 {
			subMissionId := alg.ExtractDigits(levelGroup.GroupName)
			if subMissionId == 0 || subMainMissionList[subMissionId] != nil {
				return true
			}
			return false
		}
		return mainIsLoaded
	}
	// 检查强制卸载条件
	// 检查加载条件
	// if levelGroup.GroupId == 66 {
	// 	logger.Info("")
	// 	return true
	// }
	if levelGroup.LoadCondition != nil {
		for _, conditions := range levelGroup.LoadCondition.Conditions {
			if conditions.Phase == "Finish" { // 完成了这个任务
				if finishSubMainMissionList[conditions.ID] != nil || finishMainMissionList[conditions.ID] != nil {
					isLoaded = true
					if levelGroup.LoadCondition.Operation == "Or" {
						break
					}
				} else {
					isLoaded = false
				}
				continue
			}
			if conditions.Type == "SubMission" && conditions.Phase == "" { // 接取了这个子任务
				if subMainMissionList[conditions.ID] != nil {
					isLoaded = true
					if levelGroup.LoadCondition.Operation == "Or" {
						break
					}
				} else {
					isLoaded = false
				}
				continue
			}
			if conditions.Type == "" && conditions.Phase == "" { // 接取主线任务
				if mainMissionList[conditions.ID] != nil {
					isLoaded = true
					if levelGroup.LoadCondition.Operation == "Or" {
						break
					}
				} else {
					isLoaded = false
				}
				continue
			}
		}
	}

	if !isLoaded {
		return isLoaded
	}

	// 检查卸载条件
	if levelGroup.UnloadCondition != nil {
		all := false
		for _, conditions := range levelGroup.UnloadCondition.Conditions {
			if conditions.Phase == "Finish" { // 完成了这个任务
				if finishSubMainMissionList[conditions.ID] != nil || finishMainMissionList[conditions.ID] != nil {
					if levelGroup.UnloadCondition.Operation == "Or" {
						isLoaded = false
						break
					} else {
						all = true
					}
				}
				continue
			}
			if conditions.Phase == "" { // 接取了这个任务
				if subMainMissionList[conditions.ID] != nil || mainMissionList[conditions.ID] != nil {
					if levelGroup.UnloadCondition.Operation == "Or" {
						isLoaded = false
						break
					} else {
						all = true
					}
				}
				continue
			}
		}
		if all {
			isLoaded = false
		}
	}

	return isLoaded
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
	db.blockMapLock.Lock()
	defer db.blockMapLock.Unlock()
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
	on := g.GetOnlineData()
	db := g.GetBlockMap()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
	if db[newEntryId] == nil {
		bin := database.GetBlockData(database.GSS.PlayerDataMysql,
			database.GSS.PeMysql, g.GetBasicBin().Uid, newEntryId)
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
		database.GSS.PeMysql, blockData); err != nil {
		logger.Debug("updata block data error:%s", err.Error())
	}
}

func (g *PlayerData) GetPropState(db *spb.BlockBin, groupId, propId uint32, state string) uint32 {
	if db == nil {
		return gdconf.GetStateValue(state)
	}
	on := g.GetOnlineData()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
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
	on := g.GetOnlineData()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
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
	on := g.GetOnlineData()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
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
	on := g.GetOnlineData()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
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
	on := g.GetOnlineData()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
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

func (g *PlayerData) StageObjectCapture(prop *gdconf.PropList, groupId uint32, db *spb.BlockBin) {
	if db == nil {
		return
	}
	if strings.Contains(prop.Name, "Elevator0") {
		g.ObjectCaptureUpPropState(db, groupId, prop.ID, 15)
	}
	if prop.ValueSource != nil {
		for _, v := range prop.ValueSource.Values {
			if v.Key == "IsAutoDoor" {
				g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
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

func (g *PlayerData) GetSceneAvatarByLineUP(entityGroupList *proto.SceneEntityGroupInfo, lineUp *spb.Line, leaderEntityId uint32, pos, rot *proto.Vector) {
	for sole, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar.AvatarId == 0 {
			continue
		}
		actor := &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: lineAvatar.AvatarId,
			MapLayer:     0,
			Uid:          0,
		}
		if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			conf := gdconf.GetSpecialAvatarById(lineAvatar.AvatarId)
			if conf == nil {
				continue
			}
			baseAvatarId := conf.AvatarID
			if path := gdconf.GetMultiplePathAvatarConfig(conf.AvatarID); path != nil {
				baseAvatarId = path.BaseAvatarID
			}
			actor = &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_TRIAL_TYPE,
				BaseAvatarId: baseAvatarId,
				MapLayer:     0,
				Uid:          0,
			}
		}
		entityList := &proto.SceneEntityInfo{
			EntityOneofCase: &proto.SceneEntityInfo_Actor{Actor: actor},
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
		}
		if sole == lineUp.LeaderSlot {
			entityList.EntityId = leaderEntityId
		} else {
			entityId := g.GetNextGameObjectGuid()
			entityList.EntityId = entityId
		}
		g.AddEntity(0, &AvatarEntity{
			Entity: Entity{
				EntityId: entityList.EntityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId:   actor.BaseAvatarId,
			LineAvatar: lineAvatar,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}

func (g *PlayerData) GetPropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, db *spb.BlockBin, entryId uint32) *proto.SceneEntityGroupInfo {
	for _, propList := range sceneGroup.PropList {
		entityId := g.GetNextGameObjectGuid()
		g.StageObjectCapture(propList, sceneGroup.GroupId, db)
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
					PropState: g.GetPropState(db, sceneGroup.GroupId, propList.ID, propList.State),
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
			PropId: propList.PropID,
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
			EventID: monsterList.EventID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *PlayerData) GetNPCByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) *proto.SceneEntityGroupInfo {
	for _, npcList := range sceneGroup.NPCList {
		entityId := g.GetNextGameObjectGuid()
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
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   npcList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			EntityOneofCase: &proto.SceneEntityInfo_Npc{
				Npc: &proto.SceneNpcInfo{
					ExtraInfo: nil,
					NpcId:     npcList.NPCID,
				},
			},
		}
		// 添加npc
		g.AddEntity(sceneGroup.GroupId, &NpcEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
				InstId:   npcList.ID,
			},
			NpcId: npcList.NPCID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *PlayerData) GetSceneInfo(entryId uint32, pos, rot *proto.Vector, lineUp *spb.Line) *proto.SceneInfo {
	leaderEntityId := g.GetNextGameObjectGuid()
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetServerGroup(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return nil
	}
	mazePlane := gdconf.GetMazePlaneById(mapEntrance.PlaneID)
	if mazePlane == nil {
		return nil
	}
	worldId := mazePlane.WorldID
	if worldId == 100 {
		worldId = 401
	}
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		WorldId:            worldId,
		LeaderEntityId:     leaderEntityId,
		FloorId:            mapEntrance.FloorID,
		GameModeType:       gdconf.GetPlaneType(mazePlane.PlaneType),
		PlaneId:            mapEntrance.PlaneID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		LevelGroupIdList:   make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		SceneMissionInfo:   g.GetMissionStatusBySceneInfo(gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)),
		FloorSavedData:     g.GetFloorSavedData(entryId),
		GameStoryLineId:    g.GameStoryLineId(),
		// DimensionId:        g.GetDimensionId(), // TODO
		EntityBuffList: make([]*proto.EntityBuffInfo, 0),
	}
	// scene.LightenSectionList = append(scene.LightenSectionList, 0)
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
	for _, levelGroup := range foorMap {
		if !g.IfLoadMap(levelGroup) {
			g.AddNoLoadedGroup(entryId, mapEntrance.PlaneID, mapEntrance.FloorID, levelGroup.GroupId)
			continue
		} else {
			g.AddLoadedGroup(entryId, mapEntrance.PlaneID, mapEntrance.FloorID, levelGroup.GroupId)
		}
		scene.LevelGroupIdList = append(scene.LevelGroupIdList, levelGroup.GroupId)
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
		g.GetNPCByID(entityGroupLists, levelGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}
	return scene
}

func (g *PlayerData) GetMissionStatusBySceneInfo(foorMap map[uint32]*gdconf.LevelGroup) *proto.MissionStatusBySceneInfo {
	info := &proto.MissionStatusBySceneInfo{
		DisabledMainMissionIdList:   make([]uint32, 0),
		FinishedMainMissionIdList:   make([]uint32, 0),
		SubMissionStatusList:        make([]*proto.Mission, 0),
		UnfinishedMainMissionIdList: make([]uint32, 0),
		MainMissionMcvList:          make([]*proto.MainMissionCustomValue, 0),
	}
	if foorMap == nil {
		return info
	}
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	subMissionList := g.GetSubMainMissionList()
	finishSubMissionList := g.GetFinishSubMainMissionList()
	for _, groupInfo := range foorMap {
		if groupInfo.OwnerMainMissionID != 0 {
			var isAdd = true
			for _, id := range info.DisabledMainMissionIdList {
				if id == groupInfo.OwnerMainMissionID {
					isAdd = false
					break
				}
			}
			for _, id := range info.FinishedMainMissionIdList {
				if id == groupInfo.OwnerMainMissionID {
					isAdd = false
					break
				}
			}
			if isAdd {
				var mainMissionId uint32 = 0
				if mainMissionList[groupInfo.OwnerMainMissionID] != nil {
					info.DisabledMainMissionIdList = append(info.DisabledMainMissionIdList, groupInfo.OwnerMainMissionID)
				}
				if finishMainMissionList[groupInfo.OwnerMainMissionID] != nil {
					mainMissionId = groupInfo.OwnerMainMissionID
					info.FinishedMainMissionIdList = append(info.FinishedMainMissionIdList, groupInfo.OwnerMainMissionID)
				}
				info.MainMissionMcvList = append(info.MainMissionMcvList, &proto.MainMissionCustomValue{
					MainMissionId: mainMissionId,
				})
			}
		}
		if groupInfo.AtmosphereCondition != nil {
			for _, conditions := range groupInfo.AtmosphereCondition.Conditions {
				var isAdd = true
				for _, v := range info.SubMissionStatusList {
					if v.Id == conditions.SubMissionID {
						isAdd = false
						break
					}
				}
				if isAdd {
					db := finishSubMissionList[conditions.SubMissionID]
					if db == nil {
						db = subMissionList[conditions.SubMissionID]
					}
					if db != nil {
						info.SubMissionStatusList = append(info.SubMissionStatusList, &proto.Mission{
							Status:   proto.MissionStatus(db.Status),
							Progress: db.Progress,
							Id:       conditions.SubMissionID,
						})
					} else {
						info.SubMissionStatusList = append(info.SubMissionStatusList, &proto.Mission{
							Id: conditions.SubMissionID,
						})
					}
				}
			}
		}
		if groupInfo.LoadCondition != nil {
			for _, conditions := range groupInfo.LoadCondition.Conditions {
				var isAdd = true
				if conditions.Type != "SubMission" {
					continue
				}
				for _, v := range info.SubMissionStatusList {
					if v.Id == conditions.ID {
						isAdd = false
						break
					}
				}
				if isAdd {
					db := finishSubMissionList[conditions.ID]
					if db == nil {
						db = subMissionList[conditions.ID]
					}
					if db != nil {
						info.SubMissionStatusList = append(info.SubMissionStatusList, &proto.Mission{
							Status:   proto.MissionStatus(db.Status),
							Progress: db.Progress,
							Id:       conditions.ID,
						})
					}
				}
			}
		}
		if groupInfo.UnloadCondition != nil {
			for _, conditions := range groupInfo.UnloadCondition.Conditions {
				var isAdd = true
				if conditions.Type != "SubMission" {
					continue
				}
				for _, v := range info.SubMissionStatusList {
					if v.Id == conditions.ID {
						isAdd = false
						break
					}
				}
				if isAdd {
					db := finishSubMissionList[conditions.ID]
					if db == nil {
						db = subMissionList[conditions.ID]
					}
					if db != nil {
						info.SubMissionStatusList = append(info.SubMissionStatusList, &proto.Mission{
							Status:   proto.MissionStatus(db.Status),
							Progress: db.Progress,
							Id:       conditions.ID,
						})
					}
				}
			}
		}
	}
	return info
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
			EventID: monster.EventID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 添加Npc
func (g *PlayerData) AddNpcSceneEntityRefreshInfo(mazeGroupID uint32, npcList map[uint32]*gdconf.NPCList) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, npc := range npcList {
		entityId := g.GetNextGameObjectGuid()
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
		seri := &proto.SceneEntityRefreshInfo{
			Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					GroupId:  mazeGroupID,
					InstId:   npc.ID,
					EntityId: entityId,
					Motion: &proto.MotionInfo{
						Pos: pos,
						Rot: rot,
					},
					EntityOneofCase: &proto.SceneEntityInfo_Npc{
						Npc: &proto.SceneNpcInfo{
							ExtraInfo: nil,
							NpcId:     npc.NPCID,
						},
					},
				},
			},
		}
		// 添加Npc实体
		g.AddEntity(mazeGroupID, &NpcEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  mazeGroupID,
				Pos:      pos,
				Rot:      rot,
				InstId:   npc.ID,
			},
			NpcId: npc.NPCID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 添加物品实体
func (g *PlayerData) AddPropSceneEntityRefreshInfo(mazeGroupID uint32, propList map[uint32]*gdconf.PropList, db *spb.BlockBin) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, prop := range propList {
		g.StageObjectCapture(prop, mazeGroupID, db)
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
			PropId: prop.PropID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 添加角色
func (g *PlayerData) GetAddAvatarSceneEntityRefreshInfo(lineUp *spb.Line, pos, rot *proto.Vector) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar.AvatarId == 0 {
			continue
		}
		actor := &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: lineAvatar.AvatarId,
		}
		if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			actor.AvatarType = proto.AvatarType_AVATAR_TRIAL_TYPE
		}
		entityId := g.GetNextGameObjectGuid()
		entityList := &proto.SceneEntityRefreshInfo{
			Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					EntityOneofCase: &proto.SceneEntityInfo_Actor{Actor: actor},
					Motion: &proto.MotionInfo{
						Pos: pos,
						Rot: rot,
					},
					EntityId: entityId,
				},
			},
		}
		g.AddEntity(0, &AvatarEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId:   lineAvatar.AvatarId,
			LineAvatar: lineAvatar,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, entityList)
	}
	return sceneEntityRefreshInfo
}

func (g *PlayerData) GetSceneGroupRefreshInfoByLineUP(lineUp *spb.Line, pos, rot *proto.Vector) []*proto.GroupRefreshInfo {
	groupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	for _, lineAvatar := range lineUp.AvatarIdList {
		actor := &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: lineAvatar.AvatarId,
		}
		if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			actor.AvatarType = proto.AvatarType_AVATAR_TRIAL_TYPE
		}
		entityId := g.GetNextGameObjectGuid()
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					EntityOneofCase: &proto.SceneEntityInfo_Actor{Actor: actor},
					Motion: &proto.MotionInfo{
						Pos: pos,
						Rot: rot,
					},
					EntityId: entityId,
				},
			},
		}
		g.AddEntity(0, &AvatarEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId:   lineAvatar.AvatarId,
			LineAvatar: lineAvatar,
		})
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	}
	groupRefreshInfo = append(groupRefreshInfo, sceneGroupRefreshInfo)
	return groupRefreshInfo
}

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

func (g *PlayerData) GetSpringRecoverConfig() *proto.SpringRecoverConfig {
	info := &proto.SpringRecoverConfig{
		DefaultHp:     10000,
		AutoRecoverHp: true,
		// MANPHKHEFPC: nil,
	}
	return info
}

func (g *PlayerData) GetHealPoolInfo() *proto.HealPoolInfo {
	info := &proto.HealPoolInfo{
		RefreshTime: time.Now().Unix(),
		HealPool:    23500,
	}

	return info
}
