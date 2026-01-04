package config

import "github.com/redis/go-redis/v9"

var (
	RedisClient *redis.Client
)

func InitRedisClient(rdb *redis.Client) {
	RedisClient = rdb
}