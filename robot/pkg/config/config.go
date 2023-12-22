package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel           string `json:"logLevel"`
	GameDataConfigPath string `toml:"gameDataConfigPath"`
	Server             string `json:"server"`
	AccountName        string `json:"accountName"`
	Amount             int    `json:"amount"`
	Conc               int    `json:"conc"`
	Dispatch           string `json:"dispatch"`
	GateServer         string `json:"gateServer"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

var FileNotExist = errors.New("config file not found")

func LoadConfig() error {
	filePath := "./RobotConfig.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}
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
	LogLevel:           "debug",
	GameDataConfigPath: "resources",
	Server:             "http://127.0.0.1:8080",
	AccountName:        "HkRpgRoBotGo",
	Amount:             10,
	Conc:               10,
	Dispatch:           "http://127.0.0.1:8080/query_dispatch?version=OSPRODWin1.5.0&t=114514&language_type=1&platform_type=3&channel_id=1&sub_channel_id=1&is_new_format=1",
	GateServer:         "?version=OSPRODWin1.5.0&t=114514&uid=114514&language_type=1&platform_type=3&dispatch_seed=3a57430d8d&channel_id=1&sub_channel_id=1&is_need_url=1&account_type=1&account_uid=114514",
}
