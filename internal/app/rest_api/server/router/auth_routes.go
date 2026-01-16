package router

import (
	"github.com/gin-gonic/gin"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
)

func InitializeAuthRoutes(r *gin.Engine, authClient authv1.AuthenticationServiceClient) {
	r.POST("/signup", middlewares.IsAlreadyLoggedIn(), middlewares.ValidateSignUp(), handlers.SignUp(authClient))
	r.POST("/login", middlewares.IsAlreadyLoggedIn(), middlewares.ValidateLogin(), handlers.Login(authClient))
	r.POST("/logout", middlewares.ProtectedRoute(), handlers.Logout(authClient))
}