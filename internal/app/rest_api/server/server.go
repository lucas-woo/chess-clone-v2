package server

import (
	"context"

	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server/router"
	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
	create_puzzle_client "github.com/lucas-woo/chess-clone-v2/internal/grpc/puzzles/create"
);

func CreateServer() *gin.Engine {

	middlewares.TempLogger()

	newServer := gin.New();

	newServer.Use(middlewares.Logger(), gin.Recovery())

	puzzleServer := create_puzzle_client.CreateGRPCClient()
	
	ctx := context.Background()

	router.InitializeRouter(ctx, newServer, puzzleServer)

	return newServer;
}