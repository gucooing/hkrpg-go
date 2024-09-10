package gateserver

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
)

type Config struct {
	LogLevel  string                        `json:"LogLevel"`
	MaxPlayer int64                         `json:"MaxPlayer"`
	AppList   map[string]constant.AppList   `json:"AppList"`
	NetConf   map[string]string             `json:"NetConf"`
	MysqlConf map[string]constant.MysqlConf `json:"MysqlConf"`
	RedisConf map[string]constant.RedisConf `json:"RedisConf"`
}

type NetConf struct {
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
	LogLevel:  "Info",
	MaxPlayer: -1,
	AppList: map[string]constant.AppList{
		"9001.1.1.1": {
			RegionName: "hkrpg_rel",
			MqAddr:     "127.0.0.1:20051",
			App: map[string]constant.AppNet{
				"port_player": {
					InnerAddr: "0.0.0.0",
					InnerPort: "20041",
					OuterAddr: "127.0.0.1",
					OuterPort: "20041",
				},
			},
		},
	},
	NetConf: map[string]string{
		"Node": "127.0.0.1:20081",
	},
	MysqlConf: map[string]constant.MysqlConf{
		"account": {
			Dsn: "root:password@tcp(127.0.0.1:3306)/hkrpg-go-account?charset=utf8mb4&parseTime=True&loc=Local",
		},
		"user": {
			Dsn: "root:password@tcp(127.0.0.1:3306)/hkrpg-go-user?charset=utf8mb4&parseTime=True&loc=Local",
		},
		"player": {
			Dsn: "root:password@tcp(127.0.0.1:3306)/hkrpg-go-player?charset=utf8mb4&parseTime=True&loc=Local",
		},
		"conf": {
			Dsn: "root:password@tcp(127.0.0.1:3306)/hkrpg-go-conf?charset=utf8mb4&parseTime=True&loc=Local",
		},
	},
	RedisConf: map[string]constant.RedisConf{
		"player_login": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
		"player_status": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
		"player_mail": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
		"player_chat": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
		"player_brief_data": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
	},
}
