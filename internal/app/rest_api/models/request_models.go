package models

type CreatePuzzleRequest struct {
	PlayerSide string `json:"session_id" binding:"required"`
	
}