package websocket

import "github.com/gorilla/websocket"

type Game struct { 
	socket *websocket.Conn
	gamePool *GamePool
}

func (g *Game) RunGame() {
	
}
