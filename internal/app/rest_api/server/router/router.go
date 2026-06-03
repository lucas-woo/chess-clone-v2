package router

import (
	"github.com/gin-gonic/gin"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
)

func InitializeRouter(router *gin.Engine, authClient authv1.AuthenticationServiceClient, puzzleClient puzzlesv1.PuzzlesServiceClient, profileClient usersv1.UsersServiceClient) {
	initializeTestRoutes(router, puzzleClient)
	initializeAuthRoutes(router, authClient);
	initializePuzzleRoutes(router, puzzleClient, profileClient)
	initializeAdminRoutes(router, puzzleClient)
}