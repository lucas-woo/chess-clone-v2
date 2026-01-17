package main

import (
	"fmt"
	"log"
	"net"

	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/repositories"
	authenticationgrpc "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/authentication"
	"github.com/lucas-woo/chess-clone-v2/pkg/config"
	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	redisclient "github.com/lucas-woo/chess-clone-v2/pkg/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
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

	redisClient, err := redisclient.ConnectRedis()

	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("err %v", err);
	}

	grpcServer := grpc.NewServer();

	authCollection := repositories.GetUserAuthCollection()
	authv1.RegisterAuthenticationServiceServer(grpcServer, authenticationgrpc.NewServer(redisClient, authCollection))

	healthServer := health.NewServer()
	healthv1.RegisterHealthServer(grpcServer, healthServer);
	fmt.Println("auth server running")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in starting grpc server: %v",err)
	}	
}
