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

	cursor, err := s.puzzleCollection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer cursor.Close(ctx)

	var puzzles []models.PuzzleSchema

	err = cursor.All(ctx, &puzzles)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res, err  := convertToGetPuzzleResponse(puzzles);

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return res, nil
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

func convertToGetPuzzleResponse(puzzles []models.PuzzleSchema) (*puzzlesv1.GetPuzzleResponse, error) {
	var res []*puzzlesv1.GetPuzzleResponse_Puzzle = make([]*puzzlesv1.GetPuzzleResponse_Puzzle, 0)
	for _, v := range puzzles {
		var gameState []*puzzlesv1.GetPuzzleResponse_Puzzle_PositionSchema = make([]*puzzlesv1.GetPuzzleResponse_Puzzle_PositionSchema, 0)
		for _, t := range v.GameState {
			gameState = append(gameState, &puzzlesv1.GetPuzzleResponse_Puzzle_PositionSchema{
				Piece: t.Piece,
				Placement: t.Placement,
			})
		}
		var puz *puzzlesv1.GetPuzzleResponse_Puzzle = &puzzlesv1.GetPuzzleResponse_Puzzle{
			Id: v.ID.String(),
			PlayerSide: v.PlayerSide,
			Level: v.Level,
			GameState: gameState,
			Moves: v.Moves,
		}
		res = append(res, puz)
	}
	return &puzzlesv1.GetPuzzleResponse{Puzzles: res}, nil
}

func NewServer (puzzleCollection *mongo.Collection) *Server {
	return &Server{
		puzzleCollection: puzzleCollection,
	};
}