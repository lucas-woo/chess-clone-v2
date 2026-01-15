package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
)

func SignUp(authClient authv1.AuthenticationServiceClient) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		
		newUser, err := authClient.SignUpUser(ctx.Request.Context(), &authv1.SignUpUserRequest{

		})
		if err != nil || newUser.SignupError != authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED{
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		

	}
}