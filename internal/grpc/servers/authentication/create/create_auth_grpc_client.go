package create

import (
	"log"
	"os"

	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func CreateGRPCAuthClient() authv1.AuthenticationServiceClient {
	
	port, found := os.LookupEnv("AUTH_CLIENT_PORT")

	if !found {
		log.Fatal("error with auth port env");
	}

	conn, err := grpc.NewClient("localhost:" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal()
	}

	client := authv1.NewAuthenticationServiceClient(conn)

	return client
}