package router

import (

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	
)

func InitializeRouter(router *gin.Engine, grpcClient puzzlesv1.PuzzlesServiceClient) {

	InitializeAuthRoutes(router);
}