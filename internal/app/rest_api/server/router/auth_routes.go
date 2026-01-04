package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
)

func InitializeAuthRoutes(r *gin.Engine) {
	r.POST("/signup", middlewares.IsAlreadtLoggedIn(), handlers.SignUp())
	r.POST("/login")
	r.POST("/logout")
	r.GET("verify-username");
	r.PUT("update-username");
}