package main

import (
	"fmt"
	"log"
	"net"
	"os"

	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/repositories"
	profilegrpcserver "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/users"
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

	port, found := os.LookupEnv("PROFILE_CLIENT_PORT")
	if !found {
		log.Fatal("error with profile port env");
	}	

	lis, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatalf("err %v", err);
	}

	grpcServer := grpc.NewServer();
	
	profileCollection := repositories.GetUserProfileCollection()
	usersv1.RegisterUsersServiceServer(grpcServer, profilegrpcserver.NewServer(profileCollection))
	
	healthServer := health.NewServer()
	healthv1.RegisterHealthServer(grpcServer, healthServer);
	
	fmt.Println("profile server running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in starting grpc server: %v",err)
	}
}