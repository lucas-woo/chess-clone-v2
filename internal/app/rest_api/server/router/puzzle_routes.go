package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
)

func InitializePuzzleRoutes(router *gin.Engine, puzzleClient puzzlesv1.PuzzlesServiceClient) {
	router.GET("/puzzles/get", middlewares.ProtectedRoute(), handlers.CreatePuzzleRequest(puzzleClient))
	router.POST("/puzzles/create")//admin
	router.GET("/puzzles/get-by-id" )// admin
	router.DELETE("/puzzles/delete")
}