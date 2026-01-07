package main

import (
	"fmt"
	"log"

	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	"github.com/lucas-woo/chess-clone-v2/pkg/config"
)

func main() {

	err := config.InitializeEnv();
	if err != nil {
		log.Fatal(err)
	}
	err = mongodb.ConnectMongo()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected mongo")

}
