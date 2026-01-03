package create_puzzle_client

import (
	"log"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func CreateGRPCClient() puzzlesv1.PuzzlesServiceClient {

	conn, err := grpc.NewClient("localhost:50051" ,grpc.WithTransportCredentials(insecure.NewCredentials()));

	if err != nil {
		log.Fatalf("error creating grpc client %v", err);
	}
	client := puzzlesv1.NewPuzzlesServiceClient(conn);

	return client
}