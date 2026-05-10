package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
)

func initializePuzzleRoutes(router *gin.Engine, puzzleClient puzzlesv1.PuzzlesServiceClient) {
	// use websockets later?
	// 3 min timer needs to be in sync 
	router.GET("/puzzles/", middlewares.ProtectedRoute(), handlers.CreatePuzzleRequest(puzzleClient))

	router.POST("puzzles/save")

}