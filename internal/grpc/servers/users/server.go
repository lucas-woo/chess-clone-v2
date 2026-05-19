package usersgrpc

import (
	"context"


	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"	
)

type Server struct {
	usersv1.UnimplementedUsersServiceServer
}


func (s *Server) SaveUserScore(ctx context.Context, req *usersv1.SaveUserScoreRequest) (*usersv1.SaveUserScoreResponse, error) {

	_, err := parseSaveUserScoreRequest(req);

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//check score if score is higher than current profile 

	return nil, status.Errorf(codes.Unimplemented, "method SaveUserScore not implemented")
}

func (s *Server) GetUserProfile(ctx context.Context, req *usersv1.GetUserProfileRequest) (*usersv1.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (s *Server) GetMyProfile(ctx context.Context, req *usersv1.GetMyProfileRequest) (*usersv1.GetMyProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
func (s *Server) UpdateUserProfile(ctx context.Context, req *usersv1.UpdateUserProfileRequest) (*usersv1.UpdateUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}

func parseSaveUserScoreRequest(req *usersv1.SaveUserScoreRequest) (uint32, error) {
	return req.Score, nil
}

func NewServer() *Server {
	return &Server{

	}
}