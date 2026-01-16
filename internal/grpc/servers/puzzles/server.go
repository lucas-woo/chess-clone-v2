package puzzlegrpc

import (
	"context"

	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	puzzlesv1.UnimplementedPuzzlesServiceServer
}

func (s *Server) GetPuzzle(context.Context, *puzzlesv1.GetPuzzleRequest) (*puzzlesv1.GetPuzzleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPuzzle not implemented")
}
func (s *Server) CreatePuzzle(context.Context, *puzzlesv1.CreatePuzzleRequest) (*puzzlesv1.CreatePuzzleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePuzzle not implemented")
}
func (s *Server) GetPuzzleById(context.Context, *puzzlesv1.GetPuzzleByIdRequest) (*puzzlesv1.GetPuzzleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPuzzleById not implemented")
}
func (s *Server) DeletePuzzleById(context.Context, *puzzlesv1.DeletePuzzleByIdRequest) (*puzzlesv1.DeletePuzzleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePuzzleById not implemented")
}

func NewServer () *Server {
	return &Server{

	};
}