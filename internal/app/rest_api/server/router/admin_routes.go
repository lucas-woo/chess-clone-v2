package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/handlers"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
)

func initializeAdminRoutes (router *gin.Engine, puzzleClient puzzlesv1.PuzzlesServiceClient) {


	router.POST("/admin/create-puzzle", middlewares.ProtectedAdminRoute(), handlers.CreatePuzzle(puzzleClient))


	
	router.DELETE("/admin/delete-puzzle")
	router.POST("/ban")
}