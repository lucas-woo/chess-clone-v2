package puzzlegrpc

import (
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/puzzlestore";
	"github.com/google/uuid";
	"context";
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"errors";
	"strconv"
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

	v, err := parseGetPuzzleRequest(rq);
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	puzzleArr := s.store.GetPuzzles(v);

	return convertPuzzleArrToGetPuzzleResponse(puzzleArr), nil
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

func parseGetPuzzleRequest (rq *puzzlesv1.GetPuzzleRequest) (int32, error) {

	var errs error
	val64, err := strconv.Atoi(rq.Level)

	if err != nil {
		errors.Join(err, errors.New("Couldn't parse GetPuzzleRequest"))
		return 0, errs
	}
	if val64 < 1 {
		errors.Join(err, errors.New("Puzzel level should be greater than 0 GetPuzzleRequest"))
		return 0, errs		
	}

	return int32(val64), errs
}

func convertPuzzleArrToGetPuzzleResponse(puzzleArr []*puzzlestore.Puzzle) *puzzlesv1.GetPuzzleResponse {
	var returnArray []*puzzlesv1.GetPuzzleResponse_Puzzle = make([]*puzzlesv1.GetPuzzleResponse_Puzzle, 0);

	for _, v := range puzzleArr {

		var gameStateArr []*puzzlesv1.GetPuzzleResponse_Puzzle_PositionSchema = make([]*puzzlesv1.GetPuzzleResponse_Puzzle_PositionSchema, 0);

		for _, gs := range v.GameState {
			gameStateArr = append(gameStateArr, &puzzlesv1.GetPuzzleResponse_Puzzle_PositionSchema{
				Piece: gs.Piece,
				Placement: gs.Placement,
			})
		}

		returnArray = append(returnArray, &puzzlesv1.GetPuzzleResponse_Puzzle{
			Id: v.ID.String(),
			GameState: gameStateArr,
			PlayerSide: v.PlayerSide,
			Moves: v.Moves,
			Level: v.Level,
		})
	}

	return &puzzlesv1.GetPuzzleResponse{
		Puzzles: returnArray,
	}
}

func parseCreatePuzzleRequest (rq *puzzlesv1.CreatePuzzleRequest) (*puzzlestore.Puzzle, error) {
	return nil, nil
}