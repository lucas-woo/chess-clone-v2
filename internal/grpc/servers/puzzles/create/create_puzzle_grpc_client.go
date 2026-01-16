package client

import (
	"log"
	"os"

	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func CreateGRPCPuzzleClient() puzzlesv1.PuzzlesServiceClient {

	port, found := os.LookupEnv("PUZZLE_CLIENT_PORT")
	if !found {
		log.Fatal("error with puzzle port env");
	}

	conn, err := grpc.NewClient("localhost:" + port, grpc.WithTransportCredentials(insecure.NewCredentials()));

	if err != nil {
		log.Fatalf("error creating grpc client %v", err);
	}
	client := puzzlesv1.NewPuzzlesServiceClient(conn);

	return client
}