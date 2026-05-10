package router

import (
	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
)

func initializeAdminRoutes (router *gin.Engine, puzzleClient puzzlesv1.PuzzlesServiceClient) {


	router.POST("/puzzles/create", )//admin
	router.GET("/puzzles/get-by-id" )// admin
	router.DELETE("/puzzles/delete")//admin
}