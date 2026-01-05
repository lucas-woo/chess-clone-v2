package utils

import (
	"errors"

	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
)

func ConvertCreatePuzzleRequestModel(puzzle models.CreatePuzzleRequest) (*puzzlesv1.CreatePuzzleRequest, error) {

	var err error

	if (puzzle.PlayerSide != "black" && puzzle.PlayerSide != "white") {
		err = errors.Join(err, errors.New("invalid player side"))
	}

	if puzzle.Level < 0 {
		err = errors.Join(err, errors.New("invalid puzzle level"))
	}

	if len(puzzle.Moves) <= 2 {//gotta create better validation in grpc server to not do too much computing on api server
		err = errors.Join(err, errors.New("invalid moves"))
	}

	var convertedGameState []*puzzlesv1.CreatePuzzleRequest_PositionSchema = make([]*puzzlesv1.CreatePuzzleRequest_PositionSchema, 0)
	if len(puzzle.GameState) <= 2 {
		err = errors.Join(err, errors.New("invalid game state"))
	}
	for i := 0; i < len(puzzle.GameState); i++ {
		if len(puzzle.GameState[i]) != 2 {
			err = errors.Join(err, errors.New("invalid game state"))
			break;
		}
		convertedGameState = append(convertedGameState, &puzzlesv1.CreatePuzzleRequest_PositionSchema{
			Piece: puzzle.GameState[i][0],
			Placement: puzzle.GameState[i][1],
		})
	}

	if err != nil {
		return nil, err
	}

	return &puzzlesv1.CreatePuzzleRequest{
		PlayerSide: puzzle.PlayerSide,
		Level: puzzle.Level,
		Moves: puzzle.Moves,
		GameState: convertedGameState,
	}, nil
}