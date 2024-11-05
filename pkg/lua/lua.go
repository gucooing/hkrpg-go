package lua

import (
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

var LUA *Lua

type Lua struct {
	LoginLuaList [][]byte
	PingLuaList  [][]byte
	LaLuaMap     map[string][]byte
}

type ConfLua struct {
	LoginLua []string          `json:"LoginLua"`
	PingLua  []string          `json:"PingLua"`
	LaLua    map[string]string `json:"LaLua"`
}

func LoadLua(conf *ConfLua) {
	LUA = &Lua{
		LoginLuaList: make([][]byte, 0),
		PingLuaList:  make([][]byte, 0),
		LaLuaMap:     make(map[string][]byte),
	}
	if conf == nil {
		return
	}
	n := 0
	for _, v := range conf.LoginLua {
		b, err := os.ReadFile(v)
		if err != nil {
			logger.Error(text.GetText(77), v, err.Error())
			continue
		}
		n++
		LUA.LoginLuaList = append(LUA.LoginLuaList, b)
	}
	for _, v := range conf.PingLua {
		b, err := os.ReadFile(v)
		if err != nil {
			logger.Error(text.GetText(77), v, err.Error())
			continue
		}
		n++
		LUA.PingLuaList = append(LUA.PingLuaList, b)
	}
	for k, v := range conf.LaLua {
		b, err := os.ReadFile(v)
		if err != nil {
			logger.Error(text.GetText(77), v, err.Error())
			continue
		}
		n++
		LUA.LaLuaMap[k] = b
	}
	logger.Info(text.GetText(78), n)
}

func GetLua() *Lua {
	if LUA == nil {
		LUA = &Lua{
			LoginLuaList: make([][]byte, 0),
			PingLuaList:  make([][]byte, 0),
			LaLuaMap:     make(map[string][]byte),
		}
	}
	return LUA
}

func GetLoginLua() [][]byte {
	l := GetLua()
	if l.LoginLuaList == nil {
		l.LoginLuaList = make([][]byte, 0)
	}
	return l.LoginLuaList
}

func GetPingLua() [][]byte {
	l := GetLua()
	if l.PingLuaList == nil {
		l.PingLuaList = make([][]byte, 0)
	}
	return l.PingLuaList
}

func GetLaLua(name string) []byte {
	l := GetLua()
	if l.LaLuaMap == nil {
		l.LaLuaMap = make(map[string][]byte)
		return nil
	}
	return l.LaLuaMap[name]
}
