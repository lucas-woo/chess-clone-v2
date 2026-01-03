package main

import (
	// "fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health";
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1";
	puzzlestore "github.com/lucas-woo/chess-clone-v2/internal/puzzlestore";
	puzzlegrpcserver "github.com/lucas-woo/chess-clone-v2/internal/grpc/puzzles";
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"

)

func main () {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("err %v", err);
	}

	grpcServer := grpc.NewServer();

	puzzlesv1.RegisterPuzzlesServiceServer(grpcServer, puzzlegrpcserver.NewServer(puzzlestore.NewStore()))
	healthServer := health.NewServer()
	healthv1.RegisterHealthServer(grpcServer, healthServer);

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in starting grpc server: %v",err)
	}
}