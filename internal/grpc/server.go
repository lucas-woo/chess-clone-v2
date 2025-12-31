package puzzlegrpc

import (
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/puzzlestore"
	"github.com/google/uuid"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	puzzlesv1.UnimplementedPuzzlesServiceServer
	store PuzzleStorage
}

type PuzzleStorage interface {
	CreateNewPuzzle (newPuzzle *puzzlestore.Puzzle) uuid.UUID
	GetPuzzleByID (id uuid.UUID) *puzzlestore.Puzzle
	DeletePuzzleByID (id uuid.UUID)
	GetPuzzles (level int32) []*puzzlestore.Puzzle
}


func (s *Server) GetPuzzle(_ context.Context,rq *puzzlesv1.GetPuzzleRequest) (*puzzlesv1.GetPuzzleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPuzzle not implemented")
}

func (s *Server) CreatePuzzle(_ context.Context, rq *puzzlesv1.CreatePuzzleRequest) (*puzzlesv1.CreatePuzzleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePuzzle not implemented")
}

func (s *Server) GetPuzzleById(_ context.Context, rq *puzzlesv1.GetPuzzleByIdRequest) (*puzzlesv1.GetPuzzleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPuzzleById not implemented")
}

func (s *Server) DeletePuzzleById(_ context.Context, rq *puzzlesv1.DeletePuzzleByIdRequest) (*puzzlesv1.DeletePuzzleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePuzzleById not implemented")
}