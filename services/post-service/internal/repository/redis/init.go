package redis

import (
	"github.com/1111000110/go_service/services/post-service/internal/config"
	"github.com/go-redis/redis/v8"
)

var Redisclient *redis.Client

func Init() {

}

// 选择数据库和集合
func GetClient(Password string, DB int64) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.Redis.URI, // Redis 服务器地址
		Password: Password,                   // 默认没有密码
		DB:       int(DB),                    // 默认数据库
	})
	return rdb
}
func GetDefaultCollection() *redis.Client {
	return GetClient("", 0)
}
