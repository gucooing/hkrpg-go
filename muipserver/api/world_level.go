package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
)

func worldLevel(c *gin.Context) (bool, string, bool) {
	level := alg.S2U32(c.Query("world_level"))
	if level < 0 || level > 6 {
		return false, "", false
	}
	return true, fmt.Sprintf("world_level %d", level), true
}

func getPlayerPb(c *gin.Context) (bool, string, bool) {
	return true, fmt.Sprintf("get_player_pb %s", c.Query("uid")), false
}

func status(c *gin.Context) (bool, string, bool) {
	return true, "status", false
}

// func (a *muip.Api) MaxCurAvatar(c *gin.Context) {
// 	uid := alg.S2U32(c.Query("uid"))
// 	avatarId := alg.S2U32(c.Query("avatar_id"))
// 	var isAll = false
// 	all := alg.S2U32(c.Query("all"))
// 	if all == 1 {
// 		isAll = true
// 	}
// 	message := &spb.MaxCurAvatar{
// 		PlayerUid: uid,
// 		AvatarId:  avatarId,
// 		All:       isAll,
// 	}
// 	a.ToNode(c, cmd.MaxCurAvatar, message)
// }
//
// func (a *muip.Api) GmMission(c *gin.Context) {
// 	uid := alg.S2U32(c.Query("uid"))
// 	var allFinish = false
// 	all := alg.S2U32(c.Query("all"))
// 	if all == 1 {
// 		allFinish = true
// 	}
// 	message := &spb.GmMission{
// 		PlayerUid: uid,
// 		FinishAll: allFinish,
// 		MissionId: 0,
// 	}
// 	a.ToNode(c, cmd.GmMission, message)
// }
