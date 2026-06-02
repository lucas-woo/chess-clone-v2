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
	add chan struct{}
}

func NewGame (conn *websocket.Conn, gamePool *GamePool) *Game {
	return &Game{
		socket: conn,
		gamePool: gamePool,
		currentLevel: 0,
		add: make(chan struct{}),
	}
}

func (g *Game) RunGame(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute * 3)
	defer func ()  {
		g.gamePool.leave <- g;
		g.socket.Close()
		cancel()
		//save current Level
		//let go of channels
		//need to save progress
	}()

	go func () {
		defer close(g.add)
		for {
			 _, message, err := g.socket.ReadMessage();
			 if err != nil {
				return
			 }
			 _, err = strconv.Atoi(string(message))
			 if err != nil {
				return
			 }
			 select {
			 case <-ctx.Done():
				return
			 case g.add <- struct{}{}:
			 }
		}
	}()
	
	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-g.add:
			if !ok {
				return
			}
			g.currentLevel++;
			if g.currentLevel % 5 == 0 {
				//fetch more
			}
		}

	}

}
