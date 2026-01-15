package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
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

	password, err := bcrypt.GenerateFromPassword([]byte("sdfhsdfksdhjkfshdf"), bcrypt.DefaultCost)

	res, err := client.SignUpUser(ctx, &authv1.SignUpUserRequest{
		Username: "lucas",
		Email: "1234",
		HashedPassword: string(password),
		RememberMe: false,
	})
	if res.SignupError != authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED {
		fmt.Println("error")
	}
	fmt.Println(res)



	res1, err1 := client.LoginUser(ctx, &authv1.LoginUserRequest{
		Email: "1234",
		Password: "sdfhsdfksdhjkfshdf",
	})
	if err1 != nil {
		log.Fatal("err")
	}
	if res1.LoginError != authv1.LoginUserResponse_LOGIN_ERROR_UNSPECIFIED {
		log.Fatal("invalid res\n\n")
	}

	fmt.Println(res1)
	res2, err2 := client.LogoutUser(ctx, &authv1.LogoutUserRequest{
		SessionId: res1.SessionId,
	})
	if err2 != nil {
		log.Fatal("err logout\n\n")
	}
	fmt.Println(res2)	
}