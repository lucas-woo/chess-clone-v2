package main

import (
	"log"

	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)


func main() {

	defer RedisClient.Close()
	
	if err := server.InitializeEnv(); err != nil {
		log.Fatal(err.Error())
	}
	
	rdb, err := server.ConnectRedis();
	if err != nil {
		log.Fatal(err.Error())
	}
	RedisClient = rdb
	httpServer := server.CreateServer(RedisClient)
	if err := httpServer.Run(":3000"); err != nil {
		log.Fatalf("error running client %v",err)
	}
}