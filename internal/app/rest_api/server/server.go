package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/middlewares"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/server/router"
	puzzleclient "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/puzzles/create"
	authclient "github.com/lucas-woo/chess-clone-v2/internal/grpc/servers/authentication/create"
);


func CreateServer() *gin.Engine {

	middlewares.Logger()

	newServer := gin.New();

	newServer.Use(middlewares.LoggerFunc(), gin.Recovery())

	puzzleServer := puzzleclient.CreateGRPCPuzzleClient()

	authServer := authclient.CreateGRPCAuthClient()

	router.InitializeRouter(newServer, authServer, puzzleServer)

	return newServer;
}
