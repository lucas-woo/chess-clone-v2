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
	res, err := puzzleClient.CreatePuzzle(ctx, &puzzlesv1.CreatePuzzleRequest{
		GameState: []*puzzlesv1.CreatePuzzleRequest_PositionSchema{
			{Piece: "wp", Placement: "e4"},
			{Piece: "bp", Placement: "d3"},
		},
		PlayerSide: "white",
		Moves: []string{"e4 e5","g1 f3"},
		Level: 3,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)

	req1, err := puzzleClient.GetPuzzle(ctx, &puzzlesv1.GetPuzzleRequest{
		Level: 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req1.Puzzles)
}
