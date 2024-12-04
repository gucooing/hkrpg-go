package hkrpg_go_pe

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/gateserver/session"
	"github.com/gucooing/hkrpg-go/muipserver/api"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (h *HkRpgGoServer) newHttpApi() {
	for {
		a, ok := <-h.comm.ApiChan
		if !ok {
			return
		}
		// 先查询是不是无状态指令
		if !a.IsPlayer {
			commFunc, ok := commandMap[strings.Split(a.CommandList, " ")[0]]
			if !ok {
				a.Resp <- api.ApiResp{
					Code: 404,
					Obj: gin.H{
						"code": 0,
						"msg":  "指令不存在",
					},
				}
				continue
			}
			a.Resp <- api.ApiResp{
				Code: 200,
				Obj: gin.H{
					"code": 0,
					"msg":  commFunc(h, strings.Split(a.CommandList, " ")[1:]),
				},
			}
			continue
		}
		// 玩家指令
		p := h.GetPlayer(a.Uid)
		if p == nil {
			a.Resp <- api.ApiResp{
				Code: 404,
				Obj: gin.H{
					"code": -1,
					"msg":  "Player Not Found",
				},
			}
			continue
		}

		if p.GamePlayer.RecvChan == nil {
			a.Resp <- api.ApiResp{
				Code: 404,
				Obj: gin.H{
					"code": -1,
					"msg":  "player recvchan close",
				},
			}
			continue
		}
		p.GamePlayer.ToRecvChan(player.Msg{
			Command: a.CommandList,
			MsgType: player.GmReq,
		})
		a.Resp <- api.ApiResp{
			Code: 200,
			Obj: gin.H{
				"code": 0,
				"msg":  "ok",
			},
		}
	}
}

/**********************************无状态指令*******************************/

type commHandlerFunc func(s *HkRpgGoServer, parameter []string) any

var commandMap = map[string]commHandlerFunc{
	"test":          test,
	"get_player_pb": getPlayerPb,
	"status":        status,
}

func test(s *HkRpgGoServer, parameter []string) any {
	return fmt.Sprintf("test %s", parameter)
}

func getPlayerPb(s *HkRpgGoServer, parameter []string) any {
	uid := alg.S2U32(parameter[0])
	bin := alg.S2U32(parameter[1])
	if p := s.GetPlayer(uid); p != nil {
		return p.GamePlayer.GetPd().GetBasicBin()
	} else {
		dbPlayer := database.GetPlayerDataByUid(database.PE, uid)
		if dbPlayer == nil || dbPlayer.BinData == nil {
			return "Player Not Found"
		}
		if bin == 1 {
			return hex.EncodeToString(dbPlayer.BinData)
		}
		basicBin := new(spb.PlayerBasicCompBin)
		pb.Unmarshal(dbPlayer.BinData, basicBin)
		return basicBin
	}
}

func status(s *HkRpgGoServer, parameter []string) any {
	return gin.H{
		"msg": GetServerInfo(),
	}
}

func GetServerInfo() string {
	return fmt.Sprintf("在线玩家:%v\nCPU占用:%.2f%%\n内存占用%s",
		atomic.LoadInt64(&session.CLIENT_CONN_NUM),
		alg.GetCpuOc(),
		alg.MemoryOc())
}

/**********************************分割线*******************************/

func tp(parameter []string, s *HkRpgGoServer) {
	index := len(parameter)
	if index < 3 {
		return
	}
	p := s.GetPlayer(alg.S2U32(parameter[1]))
	if p == nil {
		return
	}
	p.GamePlayer.EnterSceneByServerScNotify(alg.S2U32(parameter[2]), 0, 0, 0)
}
