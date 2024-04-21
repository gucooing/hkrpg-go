package alg

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetAppId() string {
	var appId string
	for id, sa := range os.Args {
		if sa == "-i" {
			if len(os.Args) >= id+2 {
				appId = os.Args[id+1]
			}
			break
		}
	}
	if appId == "" {
		log.Println("AppId error")
		os.Exit(0)
	}
	return appId
}

type appIdData struct {
	region  uint32
	appType uint32
	host    uint32
	index   uint32
}

func GetAppIdStr(appId uint32) string {
	data := getAppIdData(appId)
	return fmt.Sprintf("%d.%d.%d.%d", data.region, data.appType, data.host, data.index)
}

func GetAppIdUint32(appid string) uint32 {
	parts := strings.Split(appid, ".")
	if len(parts) != 4 {
		log.Println("AppId error")
		os.Exit(0)
	}
	data := appIdData{
		region:  s2u32(parts[0]),
		appType: s2u32(parts[1]),
		host:    s2u32(parts[2]),
		index:   s2u32(parts[3]),
	}
	var appID uint32
	appID |= (data.region & 0x3FF) << 22
	appID |= (data.appType & 0xF) << 18
	appID |= (data.host & 0x3FFF) << 4
	appID |= data.index & 0xF
	return appID
}

func getAppIdData(appId uint32) appIdData {
	return appIdData{
		region:  (appId >> 22) & 0x3FF,
		appType: (appId >> 18) & 0xF,
		host:    (appId >> 4) & 0x3FFF,
		index:   appId & 0xF,
	}
}

func s2u32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
