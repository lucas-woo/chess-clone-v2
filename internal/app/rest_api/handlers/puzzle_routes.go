package handlers

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
)

func GetPuzzle(puzzleClient puzzlesv1.PuzzlesServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}