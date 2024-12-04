package discord

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Type  string `json:"type"`
	Token string `json:"token"`
	Db    *Db    `json:"db"`
}
type Db struct {
	Type string `json:"type"`
	Dns  string `json:"dns"`
}

var conf *Config = nil

func getConfig() *Config {
	return conf
}

var FileNotExist = errors.New("config file not found")

func loadConfig(confName string) error {
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
	conf = c
	return nil
}
