package authenticationgrpc

import (
	"context"
	"errors"
	"time"
	// "fmt";

	"github.com/google/uuid"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/utils"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	redisclient "github.com/lucas-woo/chess-clone-v2/pkg/redis"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authv1.UnimplementedAuthenticationServiceServer; 
	redisClient *redis.Client
	userLoginCollection *mongo.Collection
	userRoleCollection *mongo.Collection
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

	_, err = s.userRoleCollection.InsertOne(ctx, &models.UserRole{
		UserID: newUser.UserID,
		Role: redisclient.UserRole,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	sessionId, err := utils.GenerateSessionId()
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}

	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, status.Error(codes.Internal, "")
	}
	createdID := id.Hex()
	
	if signupRequest.RememberMe {
		s.redisClient.Set(ctx, redisclient.SessionPrefix + sessionId, createdID, time.Second * 60 * 60 * 24)
		s.redisClient.Set(ctx, redisclient.RolePrefix + sessionId, redisclient.UserRole, time.Second * 60 * 60 * 24)
	} else {
		s.redisClient.Set(ctx, redisclient.SessionPrefix + sessionId, createdID, time.Second * 60 * 60)
		s.redisClient.Set(ctx, redisclient.RolePrefix + sessionId, redisclient.UserRole, time.Second * 60 * 60)
	}

	return &authv1.SignUpUserResponse{
		SignupError: authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED,
		SessionId: sessionId,
	}, nil
}

func (s *Server) LoginUser(ctx context.Context, loginRequest *authv1.LoginUserRequest) (*authv1.LoginUserResponse, error) {
	user, userRole, invalidInfoError, err := parseLoginUserRequest(ctx, s.userLoginCollection, s.userRoleCollection, loginRequest)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if invalidInfoError != nil {
		return &authv1.LoginUserResponse{LoginError: authv1.LoginUserResponse_LOGIN_ERROR_INVALID_CREDENTIALS}, nil
	}
	
	sessionId, err := utils.GenerateSessionId()
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}

	userID :=  user.ID.Hex()

	if loginRequest.RememberMe {
		s.redisClient.Set(ctx, redisclient.SessionPrefix + sessionId, userID, time.Second * 60 * 60 * 24)
		s.redisClient.Set(ctx, redisclient.RolePrefix + sessionId, userRole.Role, time.Second * 60 * 60 * 24)
	} else {
		s.redisClient.Set(ctx, redisclient.SessionPrefix + sessionId, userID, time.Second * 60 * 60)
		s.redisClient.Set(ctx, redisclient.RolePrefix + sessionId, userRole.Role, time.Second * 60 * 60)
	}
	return &authv1.LoginUserResponse{
		LoginError: authv1.LoginUserResponse_LOGIN_ERROR_UNSPECIFIED,
		SessionId: sessionId,
	}, nil
}

func (s *Server) LogoutUser(ctx context.Context, logoutRequest *authv1.LogoutUserRequest) (*authv1.LogoutUserResponse, error) {
	sessionID, err := parseLogoutUserRequest(logoutRequest)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	_, err = s.redisClient.Del(ctx, redisclient.SessionPrefix + sessionID).Result()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	_, err = s.redisClient.Del(ctx, redisclient.RolePrefix + sessionID).Result()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}	
	return &authv1.LogoutUserResponse{
		LoggedOut: true,
	}, nil
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
	if len(signupRequest.Password) == 0 {
		newErr := errors.New("invalid password")
		err = errors.Join(err, newErr)
	}
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)

	//gotta also double check if the username isn't taken, can't trust rest api
 	return &models.UserLogin{
		Username: signupRequest.Username,
		Hash: string(hash),
		Email: signupRequest.Email,
		UserID: uuid.New(),
	}, nil
}

func parseLoginUserRequest(ctx context.Context, userLoginCollection *mongo.Collection, userRoleCollection *mongo.Collection, loginRequest *authv1.LoginUserRequest) (*models.UserLogin, *models.UserRole, error, error) {
	//this needs validation
	var invalidInfoError error;
	if loginRequest.Password == "" {
		invalidInfoError = errors.Join(invalidInfoError, errors.New(""))
	}
	if loginRequest.Email == "" {
		invalidInfoError = errors.Join(invalidInfoError, errors.New(""))
	}

	if invalidInfoError != nil {
		return nil,nil, invalidInfoError, nil
	}

	filter := bson.D{
		bson.E{Key: "email", Value: loginRequest.Email},
	}
	var user models.UserLogin
	err := userLoginCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, nil,invalidInfoError, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(loginRequest.Password))
	if err != nil {
		return nil, nil, errors.New("invalid_credentials"), nil
	}

	roleFilter := bson.D{
		bson.E{
			Key: "uuid",
			Value: user.UserID,
		},
	}
	var userRole models.UserRole
	err = userRoleCollection.FindOne(ctx, roleFilter).Decode(&userRole)
	if err != nil {
		return nil, nil, invalidInfoError, err
	}

	return &user, &userRole, nil, nil
}

func parseLogoutUserRequest(logoutRequest *authv1.LogoutUserRequest) (string, error) {
	var err error;
	if len(logoutRequest.SessionId) < 1 {
		err = errors.New("invalid_session_id")
		return "", err
	}
	return logoutRequest.SessionId, err
}


func NewServer(redisClient *redis.Client, userLoginCollection *mongo.Collection, roleCollection *mongo.Collection) (*Server) {
	return &Server{
		redisClient: redisClient,
		userLoginCollection: userLoginCollection,
		userRoleCollection: roleCollection,
	}
}