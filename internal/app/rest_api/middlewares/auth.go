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
	return func(ctx *gin.Context) {

		var userSessionID models.UserSession;

		sessionId, err := ctx.Cookie(config.CookieSessionIDString);

		if err != nil {
			ctx.Next()
			return 
		}
		err = json.Unmarshal([]byte(sessionId), &userSessionID)
		if err != nil {
			ctx.Next()
			return 			
		}
		isValid, err := repositories.ValidateSessionID(ctx.Request.Context(), userSessionID.SessionID);
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if isValid {
			ctx.AbortWithStatus(http.StatusContinue)
			return
		}
		ctx.Next()
	}
}

func ProtectedRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var userSessionID models.UserSession;

		sessionId, err := ctx.Cookie(config.CookieSessionIDString);

		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return 
		}
		err = json.Unmarshal([]byte(sessionId), &userSessionID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return 
		}

		isValid, err := repositories.ValidateSessionID(ctx.Request.Context(), userSessionID.SessionID);
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !isValid {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return 
		}
		ctx.Next()
	}
}