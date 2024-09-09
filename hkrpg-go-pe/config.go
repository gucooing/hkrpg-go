package hkrpg_go_pe

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
)

type Config struct {
	LogLevel           string      `json:"LogLevel"`
	MaxPlayer          int64       `json:"MaxPlayer"`
	GameDataConfigPath string      `json:"GameDataConfigPath"`
	SqlPath            string      `json:"SqlPath"`
	Dispatch           *Dispatch   `json:"Dispatch"`
	GameServer         *GameServer `json:"GameServer"`
	Gm                 *Gm         `json:"Gm"`
}

type Dispatch struct {
	AutoCreate   bool            `json:"AutoCreate"`
	AppNet       constant.AppNet `json:"AppNet"`
	DispatchList []DispatchList  `json:"DispatchList"`
}
type DispatchList struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type GameServer struct {
	AppNet        constant.AppNet `json:"AppNet"`
	IsJumpMission bool            `json:"IsJumpMission"`
}
type Gm struct {
	SignKey string `json:"SignKey"`
	Addr    string `json:"Addr"`
	Port    string `json:"Port"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
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
	MaxPlayer:          -1,
	GameDataConfigPath: "resources",
	SqlPath:            "./conf/hkrpg-go-pe.db",
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
		AppNet: constant.AppNet{
			InnerAddr: "0.0.0.0",
			InnerPort: "20041",
			OuterAddr: "127.0.0.1",
			OuterPort: "20041",
		},
		IsJumpMission: true,
	},
	Gm: &Gm{
		SignKey: "123456",
		Addr:    "0.0.0.0",
		Port:    "20011",
	},
}
