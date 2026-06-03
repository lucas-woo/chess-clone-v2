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
	mu sync.Mutex
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
	defer close(g.leave)
	for game := range g.leave {
		g.mu.Lock()
		delete(g.games, game)
		g.mu.Unlock()
	}
}

func (g *GamePool) JoinGame(conn *websocket.Conn, c *gin.Context) {
	newGame := NewGame(conn, g)
	g.mu.Lock()
	g.games[newGame] = true
	g.mu.Unlock()	
	go newGame.RunGame(c)
}