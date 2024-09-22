package server

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel string    `json:"LogLevel"`
	Db       *Db       `json:"Db"`
	Host     string    `json:"Host"`
	Port     string    `json:"Port"`
	Email    *Email    `json:"Email"`
	Webhooks *Webhooks `json:"Webhooks"`
}

type Db struct {
	Type string `json:"Type"`
	Host string `json:"Host"`
}

type Email struct {
	Is       bool   `json:"Is"`
	From     string `json:"from"`
	Addr     string `json:"addr"`
	Host     string `json:"host"`
	Identity string `json:"identity"`
}

type Webhooks struct {
	Is    bool   `json:"Is"`
	Info  string `json:"Info"`
	Error string `json:"Error"`
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
	LogLevel: "info",
	Db: &Db{
		Type: "mysql",
		Host: "root:password@tcp(127.0.0.1:3306)/hkrpg-go-push?charset=utf8mb4&parseTime=True&loc=Local",
	},
	Host: "127.0.0.1",
	Port: "3000",
	Email: &Email{
		Is:       false,
		From:     "123456789@qq.com",
		Addr:     "smtp.qq.com:587",
		Host:     "smtp.qq.com",
		Identity: "123456789",
	},
	Webhooks: &Webhooks{
		Is:    false,
		Info:  "",
		Error: "",
	},
}
