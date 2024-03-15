package cmd

import (
	"reflect"
	"sync"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	pb "google.golang.org/protobuf/proto"
)

var sharedCmdProtoMap *CmdProtoMap
var cmdProtoMapOnce sync.Once

type CmdProtoMap struct {
	cmdIdProtoObjMap        map[uint16]reflect.Type
	protoObjCmdIdMap        map[reflect.Type]uint16
	cmdDeDupMap             map[uint16]bool
	cmdIdCmdNameMap         map[uint16]string
	cmdNameCmdIdMap         map[string]uint16
	cmdIdProtoObjCacheMap   map[uint16]*sync.Pool
	cmdIdProtoObjFastNewMap map[uint16]func() any
}

func GetSharedCmdProtoMap() *CmdProtoMap {
	cmdProtoMapOnce.Do(func() {
		sharedCmdProtoMap = NewCmdProtoMap()
	})
	return sharedCmdProtoMap
}

func NewCmdProtoMap() (r *CmdProtoMap) {
	r = new(CmdProtoMap)
	r.cmdIdProtoObjMap = make(map[uint16]reflect.Type)
	r.protoObjCmdIdMap = make(map[reflect.Type]uint16)
	r.cmdDeDupMap = make(map[uint16]bool)
	r.cmdIdCmdNameMap = make(map[uint16]string)
	r.cmdNameCmdIdMap = make(map[string]uint16)
	r.cmdIdProtoObjCacheMap = make(map[uint16]*sync.Pool)
	r.cmdIdProtoObjFastNewMap = make(map[uint16]func() any)
	r.registerAllMessage()

	return r
}

func (c *CmdProtoMap) regMsg(cmdId uint16, protoObjNewFunc func() any) {
	_, exist := c.cmdDeDupMap[cmdId]
	if exist {
		logger.Error("reg dup msg, cmd id: %v\n", cmdId)
		return
	} else {
		c.cmdDeDupMap[cmdId] = true
	}
	protoObj := protoObjNewFunc().(pb.Message)
	refType := reflect.TypeOf(protoObj)
	// cmdId -> protoObj
	c.cmdIdProtoObjMap[cmdId] = refType
	// protoObj -> cmdId
	c.protoObjCmdIdMap[refType] = cmdId
	cmdName := refType.Elem().Name()
	// cmdId -> cmdName
	c.cmdIdCmdNameMap[cmdId] = cmdName
	// cmdName -> cmdId
	c.cmdNameCmdIdMap[cmdName] = cmdId
	// cmdId -> protoObjCache
	c.cmdIdProtoObjCacheMap[cmdId] = &sync.Pool{
		New: protoObjNewFunc,
	}
	// cmdId -> protoObjNewFunc
	c.cmdIdProtoObjFastNewMap[cmdId] = protoObjNewFunc
}

// 性能优化专用方法 若不满足使用条件 请老老实实的用下面的反射方法

// GetProtoObjCacheByCmdId 从缓存池获取一个对象 请务必确保能容忍获取到的对象含有使用过的脏数据 否则会产生不可预料的后果
func (c *CmdProtoMap) GetProtoObjCacheByCmdId(cmdId uint16) pb.Message {
	cachePool, exist := c.cmdIdProtoObjCacheMap[cmdId]
	if !exist {
		logger.Error("unknown cmd id: %v\n", cmdId)
		return nil
	}
	protoObj := cachePool.Get().(pb.Message)
	return protoObj
}

// PutProtoObjCache 将使用结束的对象放回缓存池 请务必确保对象的生命周期真的已经结束了 否则会产生不可预料的后果
func (c *CmdProtoMap) PutProtoObjCache(cmdId uint16, protoObj pb.Message) {
	cachePool, exist := c.cmdIdProtoObjCacheMap[cmdId]
	if !exist {
		logger.Error("unknown cmd id: %v\n", cmdId)
		return
	}
	cachePool.Put(protoObj)
}

func (c *CmdProtoMap) GetProtoObjFastNewByCmdId(cmdId uint16) pb.Message {
	fn, exist := c.cmdIdProtoObjFastNewMap[cmdId]
	if !exist {
		logger.Error("unknown cmd id: %v\n", cmdId)
		return nil
	}
	protoObj := fn().(pb.Message)
	return protoObj
}

// 反射方法

func (c *CmdProtoMap) GetProtoObjByCmdId(cmdId uint16) pb.Message {
	refType, exist := c.cmdIdProtoObjMap[cmdId]
	if !exist {
		logger.Error("unknown cmd id: %v\n", cmdId)
		return nil
	}
	protoObjInst := reflect.New(refType.Elem())
	protoObj := protoObjInst.Interface().(pb.Message)
	return protoObj
}

func (c *CmdProtoMap) GetCmdIdByProtoObj(protoObj pb.Message) uint16 {
	cmdId, exist := c.protoObjCmdIdMap[reflect.TypeOf(protoObj)]
	if !exist {
		logger.Error("unknown proto object: %v\n", protoObj)
		return 0
	}
	return cmdId
}

func (c *CmdProtoMap) GetCmdNameByCmdId(cmdId uint16) string {
	cmdName, exist := c.cmdIdCmdNameMap[cmdId]
	if !exist {
		logger.Error("unknown cmd id: %v\n", cmdId)
		return ""
	}
	return cmdName
}

func (c *CmdProtoMap) GetCmdIdByCmdName(cmdName string) uint16 {
	cmdId, exist := c.cmdNameCmdIdMap[cmdName]
	if !exist {
		logger.Error("unknown cmd name: %v\n", cmdName)
		return 0
	}
	return cmdId
}
