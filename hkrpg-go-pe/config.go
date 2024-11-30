package hkrpg_go_pe

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/lua"
)

type Config struct {
	LogLevel           string      `json:"LogLevel"`
	Language           string      `json:"Language"`
	MaxPlayer          int64       `json:"MaxPlayer"`
	DataPrefix         string      `json:"DataPrefix"`
	GameDataConfigPath string      `json:"GameDataConfigPath"`
	UpstreamServerList []string    `json:"UpstreamServerList"`
	Db                 *Db         `json:"Db"`
	Dispatch           *Dispatch   `json:"Dispatch"`
	GameServer         *GameServer `json:"GameServer"`
	Gm                 *Gm         `json:"Gm"`
	PushUrl            string      `json:"PushUrl"`
}

type Db struct {
	Type string `json:"Type"`
	Dns  string `json:"Dns"`
}
type Dispatch struct {
	AutoCreate   bool              `json:"AutoCreate"`
	AppNet       constant.AppNet   `json:"AppNet"`
	DispatchList []DispatchList    `json:"DispatchList"`
	Url          *constant.UrlList `json:"Url"`
}
type DispatchList struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	DispatchUrl string `json:"dispatch_url"`
}

type GameServer struct {
	GateTcp     bool            `json:"GateTcp"`
	IsToken     bool            `json:"IsToken"`
	GetTokenUrl string          `json:"GetTokenUrl"`
	LoadLua     *lua.ConfLua    `json:"LoadLua"`
	AppNet      constant.AppNet `json:"AppNet"`
	DebugUid    uint32          `json:"DebugUid"`
	BlackCmd    map[string]bool `json:"BlackCmd"`
}
type Gm struct {
	SignKey string `json:"SignKey"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

func SetDefaultConfig() {
	CONF = DefaultConfig
}

var FileNotExist = errors.New("config file not found")

func LoadConfig(confName string) error {
	if _, err := os.Stat("./conf"); os.IsNotExist(err) {
		os.MkdirAll("./conf", 0644)
	}
	filePath := "./conf/" + confName
	f, err := os.Open(filePath)
	if err != nil {
		return FileNotExist
	}
	defer func() {
		_ = f.Close()
	}()
	c := new(Config)
	d := json.NewDecoder(f)
	if err := d.Decode(c); err != nil {
		return err
	}
	CONF = c
	return nil
}

var DefaultConfig = &Config{
	LogLevel:           "Info",
	Language:           "cn",
	MaxPlayer:          -1,
	DataPrefix:         "./data/",
	GameDataConfigPath: "resources",
	UpstreamServerList: make([]string, 0),
	Db: &Db{
		Type: "sqlite",
		Dns:  "./conf/hkrpg-go-pe.db",
	},
	PushUrl: "http://localhost:3000",
	Dispatch: &Dispatch{
		AutoCreate: true,
		AppNet: constant.AppNet{
			InnerAddr: "0.0.0.0",
			InnerPort: "8080",
			OuterAddr: "127.0.0.1",
			OuterPort: "8080",
		},
		DispatchList: []DispatchList{
			{
				Name:  "hkrpg-go",
				Title: "os_usa",
				Type:  "2",
			},
		},
	},
	GameServer: &GameServer{
		GateTcp:     false,
		IsToken:     true,
		GetTokenUrl: "http://127.0.0.1:8080",
		LoadLua: &lua.ConfLua{
			LoginLua: make([]string, 0),
			PingLua:  make([]string, 0),
			LaLua:    make(map[string]string),
		},
		AppNet: constant.AppNet{
			InnerAddr: "0.0.0.0",
			InnerPort: "20041",
			OuterAddr: "127.0.0.1",
			OuterPort: "20041",
		},
		DebugUid: 1,
		BlackCmd: map[string]bool{
			"SceneEntityMoveScRsp": true,
			"SceneEntityMoveCsReq": true,
			"PlayerHeartBeatScRsp": true,
			"PlayerHeartBeatCsReq": true,
		},
	},
	Gm: &Gm{
		SignKey: "123456",
	},
}
