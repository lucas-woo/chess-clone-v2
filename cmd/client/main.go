package main

import (
	"fmt"
	"log"

	internalconfig "github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/config"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server"
	"github.com/lucas-woo/chess-clone-v2/pkg/config"
	redisclient "github.com/lucas-woo/chess-clone-v2/pkg/redis"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)


func main() {

	if err := config.InitializeEnv(); err != nil {
		log.Fatal(err.Error())
	}
	
	rdb, err := redisclient.ConnectRedis();
	if err != nil {
		log.Fatal(err.Error())
	}

	RedisClient = rdb
	defer RedisClient.Close()

	internalconfig.InitRedisClient(RedisClient)

	httpServer := server.CreateServer()
	fmt.Println(`listening on PORT: 3000`)
	if err := httpServer.Run(":3000"); err != nil {
		log.Fatalf("error running client %v",err)
	}
}