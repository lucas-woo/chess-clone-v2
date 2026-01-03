package main

import "github.com/gin-gonic/gin";

func CreateServer() *gin.Engine {

	newServer := gin.Default()

	return newServer;
}