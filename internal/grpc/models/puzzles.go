package models

import "go.mongodb.org/mongo-driver/v2/bson"

var (
	PuzzleDatabaseName string = "puzzle_database"
	PuzzleCollectionName string = "puzzles";

	PuzzleArrayLength int = 5;
)

type PuzzleSchema struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	GameState []PuzzlePositionSchema `bson:"gameState,omitempty"`
	PlayerSide string `bson:"playerSide,omitempty"`
	Moves []string `bson:"moves,omitempty"`
	Level int32 `bson:"level,omitempty"`
}

type PuzzlePositionSchema struct {
	Piece string`bson:"piece,omitempty"`
	Placement string `bson:"placement,omitempty"`
}