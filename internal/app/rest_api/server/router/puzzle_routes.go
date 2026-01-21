package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
)

func InitializePuzzleRoutes(router *gin.Engine, puzzleClient puzzlesv1.PuzzlesServiceClient) {
	router.GET("/puzzles/get")
	router.POST("/puzzles/create")
	router.GET("/puzzles/get-by-id")
	router.DELETE("/puzzles/delete")
}