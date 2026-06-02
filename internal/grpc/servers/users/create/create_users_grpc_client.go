package create

import (
	"log"
	"os"

	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func CreateGRPCProfileClient() usersv1.UsersServiceClient {

	port, found := os.LookupEnv("PROFILE_CLIENT_PORT")
	if !found {
		log.Fatal("error with user profile port env");
	}

	conn, err := grpc.NewClient("localhost:" + port, grpc.WithTransportCredentials(insecure.NewCredentials()));

	if err != nil {
		log.Fatalf("error creating grpc client %v", err);
	}
	client := usersv1.NewUsersServiceClient(conn);

	return client
}