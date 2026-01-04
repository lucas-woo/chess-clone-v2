package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/repositories"
)

func IsAlreadtLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isValid, err := repositories.ValidateSessionID(ctx.Request.Context(), "");

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		if isValid {
			ctx.AbortWithStatus(http.StatusContinue)
		}

	}
}

func ProtectedRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}