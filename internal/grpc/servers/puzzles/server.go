package puzzlegrpc

import (
	"context"
	"errors"

	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	puzzlesv1.UnimplementedPuzzlesServiceServer
	puzzleCollection *mongo.Collection
}

func (s *Server) GetPuzzle(ctx context.Context, req *puzzlesv1.GetPuzzleRequest) (*puzzlesv1.GetPuzzleResponse, error) {

	level, err := parseGetPuzzleRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pipeline := mongo.Pipeline{
		bson.D{bson.E{Key: "$match",Value:  bson.E{Key: "level", Value: level}}}, 
		bson.D{bson.E{Key: "$sample", Value: bson.E{Key: "size", Value: models.PuzzleArrayLength}}},
	}

	s.puzzleCollection.Aggregate(ctx, pipeline)

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

func parseGetPuzzleRequest (getRequest *puzzlesv1.GetPuzzleRequest) (int32, error) {
	if getRequest.Level < 1 {
		return 0, errors.New("invalid request");
	}
	return getRequest.Level, nil
}

func NewServer (puzzleCollection *mongo.Collection) *Server {
	return &Server{
		puzzleCollection: puzzleCollection,
	};
}