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

	v, err := parseCreatePuzzleRequest(rq);

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error());
	}

	createdUUID := s.store.CreateNewPuzzle(v);

	return convertCreatePuzzleResponse(createdUUID), nil

}

func (s *Server) GetPuzzleById(_ context.Context, rq *puzzlesv1.GetPuzzleByIdRequest) (*puzzlesv1.GetPuzzleByIdResponse, error) {
	id, err := parseGetPuzzleByIdRequest(rq);
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error());
	}

	puzzle := s.store.GetPuzzleByID(id)

	return convertGetPuzzleByIdResponse(puzzle), nil
}

func (s *Server) DeletePuzzleById(_ context.Context, rq *puzzlesv1.DeletePuzzleByIdRequest) (*puzzlesv1.DeletePuzzleByIdResponse, error) {
	id, err := parseDeletePuzzleByIdRequest(rq);
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error());
	}
	s.store.DeletePuzzleByID(id);
	return &puzzlesv1.DeletePuzzleByIdResponse{
		Deleted: true,
	}, nil
}

func parseGetPuzzleRequest (rq *puzzlesv1.GetPuzzleRequest) (int32, error) {

	var errs error
	val64, err := strconv.Atoi(rq.Level)

	if err != nil {
		errs = errors.Join(errs, errors.New("Couldn't parse GetPuzzleRequest"));
		errs = errors.Join(errs, err);
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
	var errs error

	if rq.Level <= 0 { 
		errs = errors.Join(errs, errors.New("Level should be more than 0"))		
	}
	if len(rq.Moves) == 0 {
		errs = errors.Join(errs, errors.New("Moves should not be empty"))
	}
	if rq.PlayerSide != "black" && rq.PlayerSide != "white" {
		errs = errors.Join(errs, errors.New(`Player Side should either be "black" or "white"`))
	}
	if len(rq.GameState) == 0 {
		errs = errors.Join(errs, errors.New("GameState should not be empty"))
	}
	if errs != nil {
		return nil, errs
	}

	var gameStateArr []*puzzlestore.Position = make([]*puzzlestore.Position, 0);
	for _, v := range rq.GameState {
		gameStateArr = append(gameStateArr, &puzzlestore.Position{
			Piece: v.Piece,
			Placement: v.Placement,
		})
	}

	return &puzzlestore.Puzzle{
		GameState: gameStateArr,
		PlayerSide: rq.PlayerSide,
		Level: rq.Level,
		Moves: rq.Moves,
	}, nil
}

func convertCreatePuzzleResponse (newUUID uuid.UUID) *puzzlesv1.CreatePuzzleResponse {
	return &puzzlesv1.CreatePuzzleResponse{
		Id: newUUID.String(),
	}
}

func parseGetPuzzleByIdRequest (rq *puzzlesv1.GetPuzzleByIdRequest) (uuid.UUID, error) {
	parsed, err := uuid.Parse(rq.Id);
	if err != nil {
		return uuid.New(), err
	}
	return parsed, nil
}

func convertGetPuzzleByIdResponse (puzzle *puzzlestore.Puzzle) *puzzlesv1.GetPuzzleByIdResponse {
	var gameStateArr []*puzzlesv1.GetPuzzleByIdResponse_PositionSchema = make([]*puzzlesv1.GetPuzzleByIdResponse_PositionSchema, 0);
	for _, v := range puzzle.GameState {
		gameStateArr = append(gameStateArr, &puzzlesv1.GetPuzzleByIdResponse_PositionSchema{
			Piece: v.Piece,
			Placement: v.Placement,
		})
	}
	return &puzzlesv1.GetPuzzleByIdResponse{
		PlayerSide: puzzle.PlayerSide,
		GameState: gameStateArr,
		Level: puzzle.Level,
		Moves: puzzle.Moves,
	}
}

func parseDeletePuzzleByIdRequest(rq *puzzlesv1.DeletePuzzleByIdRequest) (uuid.UUID, error) {
	parsed, err := uuid.Parse(rq.Id);
	if err != nil {
		return uuid.New(), err
	}
	return parsed, nil
}

func NewServer (newStore *puzzlestore.Store) *Server {
	return &Server{
		store: newStore,
	};
}