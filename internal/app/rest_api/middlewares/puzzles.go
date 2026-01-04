package middlewares

import "github.com/gin-gonic/gin"



func ExtractPuzzle() gin.HandlerFunc {
	return func(ctx *gin.Context) {



		ctx.Next()
	}
}