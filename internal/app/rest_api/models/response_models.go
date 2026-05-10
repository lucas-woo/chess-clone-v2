package models

type CreatePuzzleResponse struct {
	Id string `json:"id" binding:"required"`
}

type DeletePuzzleResponse struct {
	Deleted bool `json:"deleted" binding:"required"`
}