package authenticationgrpc

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	redisclient "github.com/lucas-woo/chess-clone-v2/pkg/redis"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authv1.UnimplementedAuthenticationServiceServer; 
	redisClient *redis.Client
	userLoginCollection *mongo.Collection
}

func (s *Server) SignUpUser(ctx context.Context, signupRequest *authv1.SignUpUserRequest) (*authv1.SignUpUserResponse, error) {

	newUser, err := parseSignUpUserRequest(signupRequest)
	
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	result, err := s.userLoginCollection.InsertOne(ctx, newUser);

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	sessionId := uuid.New()
	createdID := result.InsertedID.(bson.ObjectID).String()

	if signupRequest.RememberMe {
		s.redisClient.Set(ctx, redisclient.SessionPrefix + sessionId.String(), createdID, time.Second * 60 * 60 * 24)
	} else {
		s.redisClient.Set(ctx, redisclient.SessionPrefix + sessionId.String(), createdID, time.Second * 60 * 60)
	}

	return &authv1.SignUpUserResponse{
		SignupError: authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED,
		SessionId: sessionId.String(),
	}, nil
}


func (s *Server) LoginUser(ctx context.Context, loginRequest *authv1.LoginUserRequest) (*authv1.LoginUserResponse, error) {

	return nil, nil
}
func (s *Server) LogoutUser(context.Context, *authv1.LogoutUserRequest) (*authv1.LogoutUserResponse, error) {
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

func parseLoginUserRequest(loginReq *authv1.LoginUserRequest) (*models.UserLogin, error) {
	
	return nil, nil
}


func NewServer(redisClient *redis.Client, userLoginCollection *mongo.Collection) (*Server) {
	return &Server{
		redisClient: redisClient,
		userLoginCollection: userLoginCollection,
	}
}