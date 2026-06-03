package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
)

func initializePuzzleRoutes(router *gin.Engine, puzzleClient puzzlesv1.PuzzlesServiceClient, profileClient usersv1.UsersServiceClient) {

	router.GET("/puzzles/play", middlewares.ProtectedRoute(), handlers.PlayPuzzle(puzzleClient, profileClient))


}