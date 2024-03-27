package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel string             `json:"LogLevel"`
	AppList  map[string]AppList `json:"AppList"`
}
type AppList struct {
	App map[string]App `json:"app"`
}
type App struct {
	Port string `json:"port"`
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
	LogLevel: "Info",
	AppList: map[string]AppList{
		"9001.1.1.1": {
			App: map[string]App{
				"port_player": {
					Port: "20041",
				},
			},
		},
		"9001.2.1.1": {
			App: map[string]App{
				"port_gt": {
					Port: "20071",
				},
			},
		},
		"9001.3.1.1": {
			App: map[string]App{
				"port_service": {
					Port: "20081",
				},
			},
		},
		"9001.4.1.1": {
			App: map[string]App{
				"port_http": {
					Port: "8080",
				},
			},
		},
		"9001.5.1.1": {
			App: map[string]App{
				"port_http": {
					Port: "20011",
				},
			},
		},
	},
}
