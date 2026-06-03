package websocket

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
	usersv1 "github.com/lucas-woo/chess-clone-v2/api/users/v1"
)

type Game struct { 
	socket *websocket.Conn
	gamePool *GamePool
	currentLevel int
	userId string
}

type GameMessage struct {
	Message string `json:"message"`
}

func NewGame (conn *websocket.Conn, gamePool *GamePool, userId string) *Game {
	return &Game{
		socket: conn,
		gamePool: gamePool,
		userId: userId,
		currentLevel: 0,
	}
}

func (g *Game) RunGame(c *gin.Context, puzzleClient puzzlesv1.PuzzlesServiceClient, profileClient usersv1.UsersServiceClient) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute * 3)

	defer cancel()

	go func() {
		<-ctx.Done()

		g.gamePool.leave <- g;
		
		err := g.socket.WriteJSON(&GameMessage{
			Message: string(g.currentLevel),
		})
		if err != nil {
			g.socket.Close()
			return 
		}

		//save in server

		g.socket.Close()

	}()

	err := g.socket.WriteJSON(&GameMessage{
		Message: "start",
	})
	if err != nil {
		return
	}

	for {
		var msg GameMessage
		err := g.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		if msg.Message != "next" {
			return
		}
		g.currentLevel++;
		if g.currentLevel % 5 == 0 {
			//send req and send client
		}

	}

}
