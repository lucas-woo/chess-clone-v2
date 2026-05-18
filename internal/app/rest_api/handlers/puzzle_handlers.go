package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/websocket"
)

func PlayPuzzle(puzzleClient puzzlesv1.PuzzlesServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		


	}
}