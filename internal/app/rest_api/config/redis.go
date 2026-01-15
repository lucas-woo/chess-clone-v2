package config

import "github.com/redis/go-redis/v9"

var (
	RedisClient *redis.Client
	CookieSessionIDString = "session_id"
)


func InitRedisClient(rdb *redis.Client) {
	RedisClient = rdb
}