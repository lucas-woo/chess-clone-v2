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

	res, err := client.GetPuzzleById(ctx, &puzzlesv1.GetPuzzleByIdRequest{

	})
	if err != nil {
		log.Fatalf("error with req %v", err);
	}

	fmt.Println(res)
}