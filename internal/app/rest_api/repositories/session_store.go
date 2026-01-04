package repositories

import (
	"context"

	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server/router"
	"github.com/redis/go-redis/v9"
)


func ValidateSessionID(ctx context.Context, sessionID string) (bool, error) {
	_, err := router.RedisClient.HGet(ctx, models.RedisSessionKey, sessionID).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err	
	}

	return true, nil
}