package alg

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	data := appIdData{
		region:  (appId >> 22) & 0x3FF,
		appType: (appId >> 18) & 0xF,
		host:    (appId >> 4) & 0x3FFF,
		index:   appId & 0xF,
	}
	return fmt.Sprintf("%d.%d.%d.%d", data.region, data.appType, data.host, data.index)
}

func GetAppIdUint32(appid string) uint32 {
	parts := strings.Split(appid, ".")
	if len(parts) != 4 {
		log.Println("AppId error")
		os.Exit(0)
	}
	data := appIdData{
		region:  S2U32(parts[0]),
		appType: S2U32(parts[1]),
		host:    S2U32(parts[2]),
		index:   S2U32(parts[3]),
	}
	var appID uint32
	appID |= (data.region & 0x3FF) << 22
	appID |= (data.appType & 0xF) << 18
	appID |= (data.host & 0x3FFF) << 4
	appID |= data.index & 0xF
	return appID
}

func S2U32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

func S2I32(msg string) int32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return int32(ms)
}

func GetEveryDay4() time.Duration {
	currentTime := time.Now()
	nextExecution := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 4, 0, 0, 0, currentTime.Location())
	if currentTime.Hour() >= 4 {
		nextExecution = nextExecution.AddDate(0, 0, 1)
	}
	return nextExecution.Sub(currentTime)
}

func GetDaysSinceBaseline(currentTime time.Time) int32 {
	baselineAt4 := time.Date(2024, 8, 1, 4, 0, 0, 0, currentTime.Location())
	baselineAt4 = baselineAt4.AddDate(0, 0, -1)
	daysSinceBaseline := int32(currentTime.Sub(baselineAt4).Hours() / 24)
	return daysSinceBaseline
}

func ExtractDigits(str string) uint32 {
	re := regexp.MustCompile(`\d+`)
	parts := re.FindAllString(str, -1)

	if len(parts) != 2 {
		return 0
	}
	firstPart := parts[0]
	secondPart := parts[1]

	var thirdPart string
	if len(secondPart) < 2 {
		thirdPart = fmt.Sprintf("%02s", secondPart)
	} else {
		thirdPart = secondPart[len(secondPart)-2:]
	}
	result := firstPart + thirdPart
	return S2U32(result)
}
