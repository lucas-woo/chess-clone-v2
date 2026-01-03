package main

import (
	"log"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server"
)


func main() {
	
	if err := server.ConnectRedis(); err != nil {
		log.Fatal(err.Error())
	}

	httpServer := server.CreateServer()
	if err := httpServer.Run(":3000"); err != nil {
		log.Fatalf("error running client %v",err)
	}
}