package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
)

func CreatePuzzle(puzzleClient puzzlesv1.PuzzlesServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		p, exists := c.Get(models.CreatedPuzzleKey)
		if !exists {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}
		puzzle, ok := p.(*puzzlesv1.CreatePuzzleRequest)
		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}

		puzzleResponse, err := puzzleClient.CreatePuzzle(c.Request.Context(), puzzle)

		if err != nil || puzzleResponse.Id == "" {
			c.AbortWithStatus(http.StatusInternalServerError)
			return 			
		}

		res := models.CreatePuzzleResponse{
			Id: puzzleResponse.Id,
		}

		c.JSON(http.StatusCreated, res)
	}
}


func DeletePuzzle(puzzleClient puzzlesv1.PuzzlesServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var puzzleID models.DeletePuzzleRequest
		if err := c.ShouldBindBodyWithJSON(&puzzleID); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		res, err := puzzleClient.DeletePuzzleById(c.Request.Context(), &puzzlesv1.DeletePuzzleByIdRequest{
			Id: puzzleID.Id,
		})
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return			
		}

		deletedResponse := models.DeletePuzzleResponse{
			Deleted: res.Deleted,
		}

		c.JSON(http.StatusOK, deletedResponse)

	}
}