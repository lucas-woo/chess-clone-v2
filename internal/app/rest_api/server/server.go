package server

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server/router"
	create_puzzle_client "github.com/lucas-woo/chess-clone-v2/internal/grpc/puzzles/create"
	"github.com/lucas-woo/godotenv"
	"github.com/redis/go-redis/v9"
);



func CreateServer() *gin.Engine {

	middlewares.Logger()

	newServer := gin.New();

	newServer.Use(middlewares.LoggerFunc(), gin.Recovery())

	puzzleServer := create_puzzle_client.CreateGRPCClient()

	router.InitializeRouter(newServer, puzzleServer)

	return newServer;
}

func ConnectRedis() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR");
	redisPassword := os.Getenv("REDIS_PASS");
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}	
	redisProtocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		return nil, err
	}	
	rdb := redis.NewClient(&redis.Options{
        Addr: redisAddr,
        Password: redisPassword, 
        DB: redisDB,  
				Protocol: redisProtocol,              
	})
	if rdb == nil {
		return nil, errors.New("error connecting to redis client")
	}
	fmt.Println("connected redis client")
	return rdb, nil
}

func InitializeEnv() error {
	return godotenv.LoadEnv()
}