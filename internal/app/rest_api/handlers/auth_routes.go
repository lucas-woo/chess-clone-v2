package handlers

import (
	"github.com/gin-gonic/gin"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
)

func SignUp(authClient authv1.AuthenticationServiceClient) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		//validation should happen here
		//signup user to db and create session 
		// authClient.SignUpUser()
	}
}