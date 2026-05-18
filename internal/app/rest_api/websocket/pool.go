package websocket

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type GamePool struct {
	mu sync.Mutex
}

var games *GamePool 

func GetGamePool() (*GamePool, error) {
	if games == nil {
		return nil, errors.New("no player pool")
	}
	return games, nil
}

func NewGamePool () *GamePool {
	games = &GamePool{}
	return games
}

func (p *GamePool) JoinGame(conn *websocket.Conn) {

}