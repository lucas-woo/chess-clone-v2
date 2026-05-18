package websocket

import (
	"context"
	"strconv"

	"github.com/gorilla/websocket"
)

type Game struct { 
	socket *websocket.Conn
	gamePool *GamePool
	currentLevel int
	add chan struct{}
}

func (g *Game) RunGame(ctx context.Context) {
	defer func ()  {
		g.gamePool.leave <- g;
		g.socket.Close()
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
