package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
)

func InitializeAuthRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignUp())
	r.POST("/login")
	r.POST("/logout")
	r.GET("verify-username");
	r.PUT("update-username");
}