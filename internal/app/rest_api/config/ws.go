package config

import "github.com/gorilla/websocket"

const (
	ReadBufferSize int = 1024
	WriteBufferSize int = 1024
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize: ReadBufferSize,
	WriteBufferSize: WriteBufferSize,
}