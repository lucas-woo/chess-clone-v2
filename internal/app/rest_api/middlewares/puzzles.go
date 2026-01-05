package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/utils"
)


func ExtractPuzzle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var puzzle models.CreatePuzzleRequest;
		
		err := ctx.ShouldBindBodyWithJSON(&puzzle);

		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		convertedPuzzle, err := utils.ConvertCreatePuzzleRequestModel(puzzle)

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(models.CreatedPuzzleKey, convertedPuzzle)
		ctx.Next()
	}
}