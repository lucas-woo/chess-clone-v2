package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/config"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/repositories"
	redisclient "github.com/lucas-woo/chess-clone-v2/pkg/redis"
)

func IsAlreadyLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie(config.CookieSessionIDString);
		
		if err != nil {
			c.Next()
			return 
		}
		isValid, err := repositories.ValidateSessionID(c.Request.Context(), sessionID);

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if isValid {
			c.AbortWithStatus(http.StatusContinue)
			return
		}
		c.Next()
	}
}

func ProtectedRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		sessionId, err := c.Cookie(config.CookieSessionIDString);

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}

		isValid, err := repositories.ValidateSessionID(c.Request.Context(), sessionId);

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !isValid {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}

		c.Set(config.CookieSessionIDString, sessionId)

		c.Next()
	}
}

//needs to validate the signup request body, username, password, email...
func ValidateSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signUpReq models.UserSignUpData
		if err := c.ShouldBindBodyWithJSON(&signUpReq); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set(config.UserSignUpData, signUpReq)
		c.Next()
	}
}

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequestBody models.UserLoginData

		if c.ShouldBindBodyWithJSON(&loginRequestBody) != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		
		c.Set(config.UserLoginData, loginRequestBody)
		c.Next()
	}
}

func ProtectedAdminRoute () gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionID, err := c.Cookie(config.CookieSessionIDString);

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}
		
		userRole, err := repositories.GetUserRole(c.Request.Context(), sessionID)
	
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return 			
		}

		if userRole != redisclient.AdminRole {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		c.Next()
	}
}