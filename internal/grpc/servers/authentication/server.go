package authenticationgrpc

import (
	"context"

	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
)

type Server struct {
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