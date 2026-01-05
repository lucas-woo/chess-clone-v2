package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
)

func CreatePuzzleRequest(puzzleGrpcCient puzzlesv1.PuzzlesServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var puzzle *puzzlesv1.CreatePuzzleRequest;
		temp, exists := ctx.Get(models.CreatedPuzzleKey)
		if !exists {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		puzzle = temp.(*puzzlesv1.CreatePuzzleRequest)
		res, err := puzzleGrpcCient.CreatePuzzle(ctx.Request.Context(), puzzle)

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return			
		}
		var response models.CreatePuzzleResponse = models.CreatePuzzleResponse{Id: res.Id}
		ctx.JSON(http.StatusCreated, response)
	}
}