package models

type CreatePuzzleRequest struct {
	PlayerSide string `json:"playerSide" binding:"required"`
	Level int32 `json:"level" binding:"required"`
	GameState [][]string `json:"gameState" binding:"required"`
	Moves []string `json:"moves" binding:"required"`
}