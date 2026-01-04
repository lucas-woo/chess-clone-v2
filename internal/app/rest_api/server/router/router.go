package router

import (
	"context"

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func InitializeRouter(ctx context.Context, router *gin.Engine, grpcClient puzzlesv1.PuzzlesServiceClient, redisClient *redis.Client) {
	
	RedisClient = redisClient

	InitializeAuthRoutes(router);
}