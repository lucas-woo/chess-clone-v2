package main

import (
	"context"
	"fmt"
	"log"

	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	client "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/users/create"
	"github.com/lucas-woo/chess-clone-v2/pkg/config"
)

func main() {
	err := config.InitializeEnv()
	if err != nil {
		log.Fatal(err)
	}
	profileClient := client.CreateGRPCProfileClient()
	ctx := context.Background()

	res, err := profileClient.SaveUserScore(ctx, &usersv1.SaveUserScoreRequest{
		Score: 4,
		UserId: "",
	})
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(res.Updated)
}
