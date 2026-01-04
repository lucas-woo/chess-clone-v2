package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
)

func InitializeTestRoutes(router *gin.Engine, puzzleGRPCclient puzzlesv1.PuzzlesServiceClient) {
	router.POST("/newPuzzle", handlers.CreatePuzzleRequest(puzzleGRPCclient))
}