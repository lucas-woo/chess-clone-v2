package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/config"
)

func SignUp(authClient authv1.AuthenticationServiceClient) gin.HandlerFunc {
	return func (c *gin.Context) {
		
		newUser, err := authClient.SignUpUser(c.Request.Context(), &authv1.SignUpUserRequest{

		})
		if err != nil || newUser.SignupError != authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.SetCookie(config.CookieSessionIDString, newUser.SessionId, config.CookieSessionMaxAge, config.CookieSessionPath, config.CookieSessionDomain, config.CookieSessionSecure, config.CookieSessionHttpOnly)

	}
}