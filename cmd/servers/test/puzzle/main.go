package main

import (
	"context"
	"fmt"
	"log"

	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	client "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/puzzles/create"
	"github.com/lucas-woo/chess-clone-v2/pkg/config"
)

func main() {
	err := config.InitializeEnv()
	if err != nil {
		log.Fatal(err)
	}
	puzzleClient := client.CreateGRPCPuzzleClient()
	ctx := context.Background()
	res, err := puzzleClient.GetPuzzle(ctx, &puzzlesv1.GetPuzzleRequest{
		Level: 2,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}
