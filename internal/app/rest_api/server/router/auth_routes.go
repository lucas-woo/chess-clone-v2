package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeAuthRoutes(r *gin.Engine) {
	r.POST("/signup", )
	r.POST("/login")
	r.POST("/logout")
	r.GET("verify-username");
	r.PUT("update-username");
}