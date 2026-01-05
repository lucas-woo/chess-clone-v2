package authenticationgrpc

import (
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
)

type Server struct {
	authv1.UnimplementedAuthenticationServiceServer; 
}