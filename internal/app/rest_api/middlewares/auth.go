package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/repositories"
)

func IsAlreadtLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repositories.ValidateSessionID(ctx.Request.Context(), "")
	}
}

func ReturnServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, "")
}


func ProtectedRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}