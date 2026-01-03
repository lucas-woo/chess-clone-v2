package router

import (
	"context"

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
)

func InitializeRouter(ctx context.Context, ginEngine *gin.Engine, grpcClient puzzlesv1.PuzzlesServiceClient) {
	
}