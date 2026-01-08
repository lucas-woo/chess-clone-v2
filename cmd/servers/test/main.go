package main

import (
	"context"
	"fmt"
	"log"

	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()));

	if err != nil {
		log.Fatalf("error creating grpc client %v", err);
	}
	client := authv1.NewAuthenticationServiceClient(conn)
	
	ctx := context.Background()
	res, err := client.SignUpUser(ctx, &authv1.SignUpUserRequest{
		Username: "lucas",
		Email: "1234",
		HashedPassword: "sdfhsdfksdhjkfshdf",
		RememberMe: false,
	})
	if res.SignupError != authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED {
		fmt.Println("error")
	}
	fmt.Println(res)
}