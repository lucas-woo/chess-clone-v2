package server

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server/router"
	create_puzzle_client "github.com/lucas-woo/chess-clone-v2/internal/grpc/puzzles/create"
	"github.com/redis/go-redis/v9"
);


var (
	redisClient *redis.Client
)

func CreateServer() *gin.Engine {

	middlewares.TempLogger()

	newServer := gin.New();

	newServer.Use(middlewares.Logger(), gin.Recovery())

	puzzleServer := create_puzzle_client.CreateGRPCClient()
	
	ctx := context.Background()

	router.InitializeRouter(ctx, newServer, puzzleServer)

	return newServer;
}

func ConnectRedis() error {
	redisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "", 
        DB: 0,                
	})
	if redisClient == nil {
		return errors.New("error connecting to redis client")
	}
	return nil
}