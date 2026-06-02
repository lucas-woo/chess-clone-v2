package websocket

import "github.com/gorilla/websocket"

const (
	readBufferSize int = 1024
	writeBufferSize int = 1024
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize: readBufferSize,
	WriteBufferSize: writeBufferSize,
}