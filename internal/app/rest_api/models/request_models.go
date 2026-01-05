package models

var (
	CreatedPuzzleKey string = "created_puzzle"
)

type CreatePuzzleRequest struct {
	PlayerSide string `json:"playerSide" binding:"required"`
	Level int32 `json:"level" binding:"required"`
	GameState [][]string `json:"gameState" binding:"required"`
	Moves []string `json:"moves" binding:"required"`
}

type CreatePuzzleResponse struct {
	Id string `json:"id" binding:"required"`
}