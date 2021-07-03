package database

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
)

/**
redis配置
*/
var redisDB *redis.Client
var once2 sync.Once

func initClient() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisDB.Ping().Result() //测试
	if err != nil {
		log.Println("redis连接初始化失败！")
		panic(err)
	}
}

func GetRedisDB() *redis.Client {
	once2.Do(initClient)
	return redisDB
}
