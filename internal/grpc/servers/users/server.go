package usersgrpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	usersv1.UnimplementedUsersServiceServer
	userProfileCollection *mongo.Collection
}


func (s *Server) SaveUserScore(ctx context.Context, req *usersv1.SaveUserScoreRequest) (*usersv1.SaveUserScoreResponse, error) {

	_, err := parseSaveUserScoreRequest(req);

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var userProfile models.UserProfile

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	filter := bson.D{
		bson.E{Key: "", Value: userId},
	}

	err = s.userProfileCollection.FindOne(ctx, filter).Decode(&userProfile)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	
	//check score if score is higher than current profile 


	return nil, status.Errorf(codes.Unimplemented, "method SaveUserScore not implemented")
}

//implement later
func (s *Server) GetUserProfile(ctx context.Context, req *usersv1.GetUserProfileRequest) (*usersv1.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
//implement later
func (s *Server) GetMyProfile(ctx context.Context, req *usersv1.GetMyProfileRequest) (*usersv1.GetMyProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
//implement later
func (s *Server) UpdateUserProfile(ctx context.Context, req *usersv1.UpdateUserProfileRequest) (*usersv1.UpdateUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}

func parseSaveUserScoreRequest(req *usersv1.SaveUserScoreRequest) (uint32, error) {
	if req.Score <= 0 {
		return 0, errors.New("")
	}
	return req.Score, nil
}

func NewServer(userProfileCollection *mongo.Collection) *Server {
	return &Server{
		userProfileCollection: userProfileCollection,
	}
}