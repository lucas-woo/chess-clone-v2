package config

import "github.com/redis/go-redis/v9"

var (
	RedisClient *redis.Client
	RedisSessionKey string = "user_session"
	CookieSessionIDString = "session_id"
)

type UserSession struct {
	SessionID string `json:"session_id" binding:"required"`
}

func InitRedisClient(rdb *redis.Client) {
	RedisClient = rdb
}