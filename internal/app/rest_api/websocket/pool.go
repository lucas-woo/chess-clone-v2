package websocket

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type GamePool struct {
	games map[*Game]bool
	leave chan *Game
}

var (
	gamePool *GamePool 
	once sync.Once
)

func GetGamePool() (*GamePool, error) {
	if gamePool == nil {
		return nil, errors.New("no player pool")
	}
	return gamePool, nil
}

func NewGamePool () *GamePool {
	once.Do(func ()  {
		gamePool = &GamePool{
			games: make(map[*Game]bool),
			leave: make(chan *Game),
		}
		go gamePool.run()		
	})
	return gamePool
}

func (g *GamePool) run() {
	for game := range g.leave {
		delete(g.games, game)
	}
}

func (p *GamePool) JoinGame(conn *websocket.Conn, ctx *gin.Context) {
	
}