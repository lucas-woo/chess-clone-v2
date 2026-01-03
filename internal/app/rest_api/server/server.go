package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
	// create_puzzle_client "github.com/lucas-woo/chess-clone-v2/internal/grpc/puzzles/create"
	// "context"
);

func CreateServer() *gin.Engine {

	middlewares.TempLogger()

	newServer := gin.New();

	newServer.Use(middlewares.Logger(), gin.Recovery())

	// puzzleServer := create_puzzle_client.CreateGRPCClient()
	// ctx := context.Background()
	// will add routes here

	return newServer;
}