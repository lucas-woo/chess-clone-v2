package models

type CreatePuzzleResponse struct {
	Id string `json:"id" binding:"required"`
}