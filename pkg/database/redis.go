package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
	"github.com/redis/go-redis/v9"
)

/**********************login*************************/

const (
	PlayerStatusMaxLockAliveTime = 600   // 玩家状态锁最大存活时间 秒
	MaxLockAliveTime             = 10000 // 单个锁的最大存活时间 毫秒
	LockRetryWaitTime            = 50    // 同步加锁重试间隔时间 毫秒
	MaxLockRetryTimes            = 2     // 同步加锁最大重试次数
)

// DistLockSync 加锁同步阻塞直到成功或超时
func LoginDistLockSync(rc *redis.Client, accountId string) bool {
	var result = false
	for i := 0; i < MaxLockRetryTimes; i++ {
		var err error = nil
		key := "player_login_lock:" + accountId
		result, err = rc.SetNX(ctx,
			key,
			time.Now().UnixMilli(),
			time.Millisecond*time.Duration(MaxLockAliveTime)).Result()
		if err != nil {
			log.Printf("redis lock setnx error: %v", err)
			return false
		}
		if result == true {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(LockRetryWaitTime))
	}
	return result
}

// DistUnlock 解锁
func LoginDistUnlock(rc *redis.Client, accountId string) {
	var result int64 = 0
	var err error = nil
	key := "player_login_lock:" + accountId
	result, err = rc.Del(ctx, key).Result()
	if err != nil {
		log.Printf("redis lock del error: %v", err)
		return
	}
	if result == 0 {
		log.Printf("redis lock del result is fail")
		return
	}
}

/**********************status*************************/

// 获取玩家状态
func GetPlayerStatus(rc *redis.Client, uid uint32) ([]byte, bool) {
	key := fmt.Sprintf("player_status_lock:%v", uid)
	bin, err := rc.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 标记玩家状态
func AddPlayerStatus(rc *redis.Client, uid uint32, value []byte) bool {
	var err error = nil
	key := fmt.Sprintf("player_status_lock:%v", uid)
	err = rc.Set(ctx,
		key,
		value,
		time.Second*time.Duration(PlayerStatusMaxLockAliveTime)).Err()
	if err != nil {
		logger.Error("redis lock setnx error: %v", err)
		return false
	} else {
		return true
	}
}

// 删除玩家状态
func DelPlayerStatus(rc *redis.Client, uid uint32) {
	var result int64 = 0
	var err error = nil
	key := fmt.Sprintf("player_status_lock:%v", uid)
	result, err = rc.Del(ctx, key).Result()
	if err != nil {
		logger.Error("redis lock del error: %v", err)
		return
	}
	if result == 0 {
		logger.Error("redis lock del result is fail")
		return
	}
}

/**********************player basic*************************/

func updatePlayerBasicRedis(rc *redis.Client, uid uint32, value []byte) bool {
	key := "player_brief_data:" + strconv.Itoa(int(uid))
	err := rc.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}

func getPlayerBasicRedis(rc *redis.Client, uid uint32) ([]byte, bool) {
	key := "player_brief_data:" + strconv.Itoa(int(uid))
	bin, err := rc.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

/**********************game*************************/

// 获取玩家好友申请信息
func getPlayerFriendRedis(rc *redis.Client, uid uint32) ([]byte, bool) {
	key := "player_friend:" + strconv.Itoa(int(uid))
	bin, err := rc.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 设置玩家好友申请信息
func setPlayerFriendRedis(rc *redis.Client, uid uint32, value []byte) bool {
	key := "player_friend:" + strconv.Itoa(int(uid))
	err := rc.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}

// 获取玩家待加入数据库好友信息
func getAcceptApplyFriendRedis(rc *redis.Client, uid uint32) ([]byte, bool) {
	key := "accept_apply_friend:" + strconv.Itoa(int(uid))
	bin, err := rc.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 设置玩家待加入数据库好友信息
func setAcceptApplyFriend(rc *redis.Client, uid, applyUid uint32) bool {
	key := "accept_apply_friend:" + strconv.Itoa(int(uid))
	err := rc.HSet(ctx, key, applyUid).Err()
	if err != nil {
		return false
	}
	return true
}

// 删除玩家待加入数据库好友信息
func delAcceptApplyFriendRedis(rc *redis.Client, uid uint32) {
	key := "accept_apply_friend:" + strconv.Itoa(int(uid))
	rc.Del(ctx, key)
}

// 获取玩家邮件数据
func getAllPlayerMailRedis(rc *redis.Client, uid uint32) []*constant.PlayerMail {
	var playerMail []*constant.PlayerMail
	key := "player_mail:" + strconv.Itoa(int(uid))
	db := rc.HGetAll(ctx, key).Val()
	for _, v := range db {
		mail := &constant.PlayerMail{}
		err := hjson.Unmarshal([]byte(v), &mail)
		if err != nil {
			logger.Error("get player mail error: %v", err)
			continue
		}
		playerMail = append(playerMail, mail)
	}
	return playerMail
}
