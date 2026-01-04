package handlers

import "github.com/gin-gonic/gin"

func SignUp() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		//signup user to db and create session 
	}
}