package server

import (
	"context"
	"errors"
	"os";
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server/router"
	create_puzzle_client "github.com/lucas-woo/chess-clone-v2/internal/grpc/puzzles/create"
	"github.com/lucas-woo/godotenv"
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
	redisAddr := os.Getenv("REDIS_ADDR");
	redisPassword := os.Getenv("REDIS_PASS");
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return err
	}	
	redisProtocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		return err
	}	
	redisClient = redis.NewClient(&redis.Options{
        Addr: redisAddr,
        Password: redisPassword, 
        DB: redisDB,  
				Protocol: redisProtocol,              
	})
	if redisClient == nil {
		return errors.New("error connecting to redis client")
	}
	return nil
}

func InitializeEnv() error {
	return godotenv.LoadEnv()
}