package authenticationgrpc

import (
	"context"

	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	redisClient *redis.Client
	userLoginCollection *mongo.Collection
	userProfileCollection *mongo.Collection
	authv1.UnimplementedAuthenticationServiceServer; 
}

func (s *Server) SignUpUser(context.Context, *authv1.SignUpUserRequest) (*authv1.SignUpUserResponse, error) {
	
	return nil, nil
}
func (s *Server) LoginUser(context.Context, *authv1.LoginUserRequest) (*authv1.LoginUserResponse, error) {
	return nil, nil
}
func (s *Server) LogoutUser(context.Context, *authv1.LogoutUserRequest) (*authv1.LogoutUserResponse, error) {
	return nil, nil
}
func (s *Server) VerifyUsername(context.Context, *authv1.VerifyUsernameRequest) (*authv1.VerifyUsernameResponse, error) {
	return nil, nil
}

func NewServer(redisClient *redis.Client, userLoginCollection *mongo.Collection, userProfileCollection *mongo.Collection) (*Server) {
	return &Server{
		redisClient: redisClient,
		userLoginCollection: userLoginCollection,
		userProfileCollection: userProfileCollection,
	}
}