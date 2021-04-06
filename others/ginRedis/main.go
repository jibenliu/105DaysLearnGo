package main

import (
	"context"
	"encoding/json"
	"goProjects/daylearning/others/ginRedis/engine"
	"log"
	"time"
)

var (
	ctx         = context.Background()
	redisHelper *engine.RedisHelper
)

type UserInfo struct {
	UserId string
}

func (u *UserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func main() {
	rdb := engine.GetRedisHelper()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err.Error())
	}
	userInfo := &UserInfo{UserId: "56867897283718"}
	_, err := redisHelper.Set(ctx, "moose-go", userInfo, 10*time.Minute).Result()
	if err != nil {
		log.Fatal(err)
	}

	name, err := redisHelper.Get(ctx, "moose-go").Result()
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(name), &userInfo)
}
