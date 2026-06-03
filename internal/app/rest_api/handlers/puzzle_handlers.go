package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/config"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/websocket"
)

func PlayPuzzle(puzzleClient puzzlesv1.PuzzlesServiceClient, profileClient usersv1.UsersServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		s, found := c.Get(config.CookieUserId)
		if !found {
			c.AbortWithStatus(http.StatusInternalServerError)
			return						
		}
		userId, ok := s.(string)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return								
		}
		conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		playerPool, err := websocket.GetGamePool();

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return			
		}
		playerPool.JoinGame(conn, c, puzzleClient, profileClient, userId)

	}
}