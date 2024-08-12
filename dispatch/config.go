package dispatch

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
)

type Config struct {
	LogLevel   string                        `json:"LogLevel"`
	AutoCreate bool                          `json:"AutoCreate"`
	Dispatch   []Dispatch                    `json:"Dispatch"`
	AppList    map[string]constant.AppList   `json:"AppList"`
	NetConf    map[string]string             `json:"NetConf"`
	Email      *email                        `json:"Email"`
	MysqlConf  map[string]constant.MysqlConf `json:"MysqlConf"`
	RedisConf  map[string]constant.RedisConf `json:"RedisConf"`
}
type Dispatch struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	DispatchUrl string `json:"dispatchUrl"`
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
	Dispatch: []Dispatch{
		{
			Name:        "hkrpg-go",
			Title:       "os_usa",
			Type:        "2",
			DispatchUrl: "/query_gateway",
		},
		{
			Name:        "hkrpg-official_os",
			Title:       "os_usa",
			Type:        "2",
			DispatchUrl: "/query_gateway_capture",
		},
		{
			Name:        "hkrpg-official_cn",
			Title:       "os_usa",
			Type:        "2",
			DispatchUrl: "/query_gateway_capture_cn",
		},
	},
	AppList: map[string]constant.AppList{
		"9001.1.1.1": {
			App: map[string]constant.App{
				"port_player": {
					Port: "20041",
				},
			},
		},
		"9001.2.1.1": {
			App: map[string]constant.App{
				"port_gt": {
					Port: "20071",
				},
			},
		},
		"9001.3.1.1": {
			App: map[string]constant.App{
				"port_service": {
					Port: "20081",
				},
			},
		},
		"9001.4.1.1": {
			App: map[string]constant.App{
				"port_http": {
					Port:      "8080",
					InnerAddr: "127.0.0.1",
					OuterAddr: "http://127.0.0.1:8080",
				},
			},
		},
		"9001.5.1.1": {
			App: map[string]constant.App{
				"port_service": {
					Port: "20091",
				},
			},
		},
		"9001.6.1.1": {
			App: map[string]constant.App{
				"port_http": {
					Port: "20011",
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
	Email: &email{
		From:     "123456789@qq.com",
		Addr:     "smtp.qq.com:587",
		Host:     "smtp.qq.com",
		Identity: "123456789",
	},
}
