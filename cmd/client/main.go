package main

import (
	"log"

	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server"
	"github.com/redis/go-redis/v9"
)

var (
	redisInstance *redis.Client
)

func main() {
	server := server.CreateServer()
	err := server.Run(":3000");
	if err != nil {
		log.Fatalf("error running client %v",err)
	}
}