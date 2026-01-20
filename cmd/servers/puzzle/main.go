package main

import (
	"fmt"
	"log"
	"net"
	"os"

	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/repositories"
	puzzlegrpcserver "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/puzzles"
	"github.com/lucas-woo/chess-clone-v2/pkg/config"
	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
)

func main () {

	err := config.InitializeEnv();
	if err != nil {
		log.Fatal(err)
	}
	err = mongodb.ConnectMongo()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected mongo")

	port, found := os.LookupEnv("PUZZLE_CLIENT_PORT")
	if !found {
		log.Fatal("error with puzzle port env");
	}	

	lis, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatalf("err %v", err);
	}

	grpcServer := grpc.NewServer();
	
	puzzleCollection := repositories.GetPuzzleCollection()
	puzzlesv1.RegisterPuzzlesServiceServer(grpcServer, puzzlegrpcserver.NewServer(puzzleCollection))
	
	healthServer := health.NewServer()
	healthv1.RegisterHealthServer(grpcServer, healthServer);
	fmt.Println("puzzle server running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in starting grpc server: %v",err)
	}
}