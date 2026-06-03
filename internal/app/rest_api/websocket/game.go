package websocket

import (
	"context"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Game struct { 
	socket *websocket.Conn
	gamePool *GamePool
	currentLevel int
	userId string
}

func NewGame (conn *websocket.Conn, gamePool *GamePool, userId string) *Game {
	return &Game{
		socket: conn,
		gamePool: gamePool,
		userId: userId,
		currentLevel: 0,
	}
}

func (g *Game) RunGame(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute * 3)

	defer cancel()

	go func() {
		<- ctx.Done()
		g.gamePool.leave <- g;

		err := g.socket.WriteMessage(websocket.TextMessage, []byte("start"))
		if err != nil {
			return
		}		

	}()

	err := g.socket.WriteMessage(websocket.TextMessage, []byte("start"))
	if err != nil {
		return
	}

	for {
		
	}

}
