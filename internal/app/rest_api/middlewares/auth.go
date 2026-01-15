package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/config"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/repositories"
)

func IsAlreadyLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userSessionID models.UserSession;

		sessionId, err := c.Cookie(config.CookieSessionIDString);

		if err != nil {
			c.Next()
			return 
		}
		err = json.Unmarshal([]byte(sessionId), &userSessionID)
		if err != nil {
			c.Next()
			return 			
		}
		isValid, err := repositories.ValidateSessionID(c.Request.Context(), userSessionID.SessionID);
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
		
		var userSessionID models.UserSession;

		sessionId, err := c.Cookie(config.CookieSessionIDString);

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}
		err = json.Unmarshal([]byte(sessionId), &userSessionID)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}

		isValid, err := repositories.ValidateSessionID(c.Request.Context(), userSessionID.SessionID);
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !isValid {
			c.AbortWithStatus(http.StatusBadRequest)
			return 
		}
		c.Next()
	}
}