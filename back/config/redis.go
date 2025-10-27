package config

import (
	"exchange_app/global"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
        Password: "1234", 
        DB:       0,  
	})

	global.Rdb = redisDb
}