package hkrpg_go_pe

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel           string      `json:"LogLevel"`
	GameDataConfigPath string      `json:"GameDataConfigPath"`
	SqlPath            string      `json:"SqlPath"`
	Dispatch           *Dispatch   `json:"Dispatch"`
	GameServer         *GameServer `json:"GameServer"`
	Gm                 *Gm         `json:"Gm"`
}

type Dispatch struct {
	AutoCreate   bool           `json:"AutoCreate"`
	InnerAddr    string         `json:"InnerAddr"`
	InnerPort    string         `json:"InnerPort"`
	OuterAddr    string         `json:"OuterAddr"`
	OuterPort    string         `json:"OuterPort"`
	DispatchList []DispatchList `json:"DispatchList"`
}
type DispatchList struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type GameServer struct {
	MaxPlayer     int32  `json:"MaxPlayer"`
	InnerAddr     string `json:"InnerAddr"`
	InnerPort     string `json:"InnerPort"`
	OuterAddr     string `json:"OuterAddr"`
	OuterPort     string `json:"OuterPort"`
	IsJumpMission bool   `json:"IsJumpMission"`
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
	GameDataConfigPath: "resources",
	SqlPath:            "./conf/hkrpg-go-pe.db",
	Dispatch: &Dispatch{
		AutoCreate: true,
		InnerAddr:  "0.0.0.0",
		InnerPort:  "8080",
		OuterAddr:  "127.0.0.1",
		OuterPort:  "8080",
		DispatchList: []DispatchList{
			{
				Name:  "hkrpg-go",
				Title: "os_usa",
				Type:  "2",
			},
		},
	},
	GameServer: &GameServer{
		MaxPlayer:     -1,
		InnerAddr:     "0.0.0.0",
		InnerPort:     "20041",
		OuterAddr:     "127.0.0.1",
		OuterPort:     "20041",
		IsJumpMission: true,
	},
	Gm: &Gm{
		SignKey: "123456",
		Addr:    "0.0.0.0",
		Port:    "20011",
	},
}
