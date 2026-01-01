package main

import (
	"context"
	"fmt"
	"log"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	conn, err := grpc.NewClient("localhost:50051" ,grpc.WithTransportCredentials(insecure.NewCredentials()));
	if err != nil {
		log.Fatalf("error with client %v", err);
	}

	client := puzzlesv1.NewPuzzlesServiceClient(conn);

	ctx := context.Background();

	v1, err := client.GetPuzzleById(ctx, &puzzlesv1.GetPuzzleByIdRequest{

	})
	if err != nil {
		fmt.Printf("err getting pbuid: %v",err)
	} else {
		fmt.Println(v1)
	}

	v2, err2 := client.CreatePuzzle(ctx, &puzzlesv1.CreatePuzzleRequest{

	})
	if err2 != nil {
		fmt.Printf("err: %v",err2)
	} else {
		fmt.Println(v2)
	}
	
	v3, err3 := client.CreatePuzzle(ctx, &puzzlesv1.CreatePuzzleRequest{
		PlayerSide: "white",
		Moves: []string{"e2 e4", "e7 e5"},
		Level: 3,
		GameState: []*puzzlesv1.CreatePuzzleRequest_PositionSchema{&puzzlesv1.CreatePuzzleRequest_PositionSchema{
			Piece: "wp",
			Placement: "e2",
		}},
	})
	if err3 != nil {
		fmt.Printf("err: %v",err3)
	} else {
		fmt.Println(v3)
	}	
}