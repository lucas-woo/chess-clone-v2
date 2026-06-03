package websocket

import (
	"context"
	"time"
	"strconv"
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
	Content any
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

		defer func() {
			g.gamePool.leave <- g;
			g.socket.Close()			
		}()
		
		err := g.socket.WriteJSON(&GameMessage{
			Message: "end",
			Content: strconv.Itoa(g.currentLevel),
		})
		if err != nil {
			g.socket.Close()
			return 
		}

		newCtx := c.Request.Context()
		profileClient.SaveUserScore(newCtx, &usersv1.SaveUserScoreRequest{
			Score: uint32(g.currentLevel),
			UserId: g.userId,
		})
	}()

	err := g.socket.WriteJSON(GameMessage{
		Message: "start",
	})
	if err != nil {
		return
	}

	newCtx := c.Request.Context()
	res, err := puzzleClient.GetPuzzle(newCtx, &puzzlesv1.GetPuzzleRequest{
		Level: 1,
	})
	if err != nil {
		return
	}
	err = g.socket.WriteJSON(GameMessage{
		Message: "puzzles",
		Content: res,
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
			newCtx := c.Request.Context()
			res, err := puzzleClient.GetPuzzle(newCtx, &puzzlesv1.GetPuzzleRequest{
				Level: int32((g.currentLevel + 5) / 5),
			})
			if err != nil {
				return
			}
			err = g.socket.WriteJSON(GameMessage{
				Message: "puzzles",
				Content: res,
			})
			if err != nil {
				return
			}
		}
	}
}
