package authenticationgrpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	redisClient *redis.Client
	userLoginCollection *mongo.Collection
	authv1.UnimplementedAuthenticationServiceServer; 
}

func (s *Server) SignUpUser(ctx context.Context, signupRequest *authv1.SignUpUserRequest) (*authv1.SignUpUserResponse, error) {
	
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


func parseSignUpUserRequest(signupRequest *authv1.SignUpUserRequest) (*models.UserLogin, error) {
	var err error
	//gotta add validation later, this isn't good
	if len(signupRequest.Email) == 0 {
		newErr := errors.New("invalid email")
		err = errors.Join(err, newErr)
	}
	if len(signupRequest.Username) == 0 {
		newErr := errors.New("invalid username")
		err = errors.Join(err, newErr)		
	}
	if len(signupRequest.HashedPassword) == 0 {
		newErr := errors.New("invalid password")
		err = errors.Join(err, newErr)
	}
	if err != nil {
		return nil, err
	}
	//gotta also double check if the username isn't taken, can't trust rest api
 	return &models.UserLogin{
		Username: signupRequest.Username,
		Hash: signupRequest.HashedPassword,
		Email: signupRequest.Email,
		UserID: uuid.New(),
	}, nil
}



func NewServer(redisClient *redis.Client, userLoginCollection *mongo.Collection) (*Server) {
	return &Server{
		redisClient: redisClient,
		userLoginCollection: userLoginCollection,
	}
}