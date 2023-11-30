package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel           string `json:"logLevel"`
	GameDataConfigPath string `toml:"GameDataConfigPath"` // 配置表路径
	MysqlDsn           string `json:"MysqlDsn"`
	Account            *Account
	Http               *Http
	Dispatch           []Dispatch
	Game               *Game
	Email              *email
}
type Account struct {
	AutoCreate bool  `json:"autoCreate"`
	MaxPlayer  int64 `json:"maxPlayer"`
}
type Dispatch struct {
	Name        string
	Title       string
	Type        string
	DispatchUrl string
}
type Http struct {
	Addr string `json:"addr"`
	Port int64  `json:"port"`
}
type Game struct {
	Addr string `json:"addr"`
	Port uint32 `json:"port"`
}
type email struct {
	From     string `json:"From"`
	Addr     string `json:"Addr"`
	Host     string `json:"Host"`
	Identity string `json:"Identity"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

var FileNotExist = errors.New("config file not found")

func LoadConfig() error {
	filePath := "./config.json"
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
	LogLevel:           "Info",
	GameDataConfigPath: "resources",
	MysqlDsn:           "root:password@tcp(127.0.0.1:3306)/hkrpg-go?charset=utf8mb4&parseTime=True&loc=Local",
	Account: &Account{
		AutoCreate: true,
		MaxPlayer:  -1,
	},
	Http: &Http{
		Addr: "0.0.0.0",
		Port: 8080,
	},
	Dispatch: []Dispatch{
		{
			Name:        "hkrpg-go",
			Title:       "os_usa",
			Type:        "2",
			DispatchUrl: "http://127.0.0.1:8080/query_gateway",
		},
	},
	Game: &Game{
		Addr: "127.0.0.1",
		Port: 22102,
	},
	Email: &email{
		From:     "123456789@qq.com",
		Addr:     "smtp.qq.com:587",
		Host:     "smtp.qq.com",
		Identity: "123456789",
	},
}
