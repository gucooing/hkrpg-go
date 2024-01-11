package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/random"
)

type Config struct {
	LogLevel           string       `json:"LogLevel"`
	GameDataConfigPath string       `toml:"GameDataConfigPath"`
	MysqlDsn           string       `json:"MysqlDsn"`
	Account            *Account     `json:"Account"`
	Http               *Http        `json:"Http"`
	Dispatch           []Dispatch   `json:"Dispatch"`
	Game               *Game        `json:"Game"`
	GmKey              string       `json:"GmKey"`
	Email              *email       `json:"Email"`
	Ec2b               *random.Ec2b `json:"Ec2B"`
}
type Account struct {
	AutoCreate bool  `json:"autoCreate"`
	MaxPlayer  int32 `json:"maxPlayer"`
}
type Dispatch struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	DispatchUrl string `json:"dispatchUrl"`
}
type Http struct {
	Addr        string `json:"addr"`
	Port        int64  `json:"port"`
	EnableHttps bool   `json:"enable"`
	CertFile    string `json:"certFile"`
	KeyFile     string `json:"keyFile"`
}
type Game struct {
	Addr string `json:"addr"`
	Port uint32 `json:"port"`
}
type email struct {
	From     string `json:"from"`
	Addr     string `json:"addr"`
	Host     string `json:"host"`
	Identity string `json:"identity"`
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
		Addr:        "0.0.0.0",
		Port:        8080,
		EnableHttps: true,
		CertFile:    "data/localhost.crt",
		KeyFile:     "data/localhost.key",
	},
	Dispatch: []Dispatch{
		{
			Name:        "hkrpg-go",
			Title:       "os_usa",
			Type:        "2",
			DispatchUrl: "http://127.0.0.1:8080/query_gateway",
		},
		{
			Name:        "hkrpg-official",
			Title:       "os_usa",
			Type:        "2",
			DispatchUrl: "http://127.0.0.1:8080/query_gateway_capture",
		},
	},
	Game: &Game{
		Addr: "127.0.0.1",
		Port: 22102,
	},
	GmKey: "",
	Email: &email{
		From:     "123456789@qq.com",
		Addr:     "smtp.qq.com:587",
		Host:     "smtp.qq.com",
		Identity: "123456789",
	},
}
