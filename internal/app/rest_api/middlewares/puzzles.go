package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/utils"
)


func ExtractPuzzle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var puzzle models.CreatePuzzleRequest;
		
		err := c.ShouldBindBodyWithJSON(&puzzle);

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		convertedPuzzle, err := utils.ConvertCreatePuzzleRequestModel(puzzle)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Set(models.CreatedPuzzleKey, convertedPuzzle)
		c.Next()
	}
}