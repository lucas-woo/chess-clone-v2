package main

import (
	"log"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server"
)

func main() {
	server := server.CreateServer()
	err := server.Run(":3000");
	if err != nil {
		log.Fatalf("error running client %v",err)
	}
}