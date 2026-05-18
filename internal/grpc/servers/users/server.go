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


func (s *Server) GetUserProfile(context.Context, *usersv1.GetUserProfileRequest) (*usersv1.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (s *Server) GetMyProfile(context.Context, *usersv1.GetMyProfileRequest) (*usersv1.GetMyProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
func (s *Server) UpdateUserProfile(context.Context, *usersv1.UpdateUserProfileRequest) (*usersv1.UpdateUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (s *Server) SaveUserScore(context.Context, *usersv1.SaveUserScoreRequest) (*usersv1.SaveUserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserScore not implemented")
}


func NewServer() *Server {
	return &Server{

	}
}