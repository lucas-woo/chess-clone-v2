package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authv1 "github.com/lucas-woo/chess-clone-v2/api/authentication/v1"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/config"
	"github.com/lucas-woo/chess-clone-v2/internal/app/rest_api/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(authClient authv1.AuthenticationServiceClient) gin.HandlerFunc {
	return func (c *gin.Context) {

		data, exists := c.Get(config.UserSignUpData)
		
		if !exists {
			c.AbortWithStatus(http.StatusBadRequest)
			return			
		}
		signUpUserData, ok := data.(models.UserSignUpData)

		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return						
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpUserData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return				
		}

		newUser, err := authClient.SignUpUser(c.Request.Context(), &authv1.SignUpUserRequest{
			Username: signUpUserData.Username,
			HashedPassword: string(hashedPassword),
			Email: signUpUserData.Email,
			RememberMe: signUpUserData.RememberMe,
		})

		if err != nil || newUser.SignupError != authv1.SignUpUserResponse_SIGN_UP_ERROR_UNSPECIFIED {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.SetCookie(config.CookieSessionIDString, newUser.SessionId, config.CookieSessionMaxAge, config.CookieSessionPath, config.CookieSessionDomain, config.CookieSessionSecure, config.CookieSessionHttpOnly)

		c.JSON(http.StatusCreated, "created")
	}
}

func Login(authClient authv1.AuthenticationServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		email, exists := c.Get("email")
		emailConv, ok := email.(string)
		if !exists || !ok{
			c.AbortWithStatus(http.StatusBadRequest)
			return			
		}				

		password, exists := c.Get("password")
		passwordConv, ok := password.(string)
		if !exists || !ok{
			c.AbortWithStatus(http.StatusBadRequest)
			return			
		}		

		loggedIn, err := authClient.LoginUser(c.Request.Context(), &authv1.LoginUserRequest{
			Email: emailConv,
			Password: passwordConv,
		})
		
		if err != nil || loggedIn.LoginError != authv1.LoginUserResponse_LOGIN_ERROR_UNSPECIFIED {
			c.AbortWithStatus(http.StatusBadRequest)
			return						
		}

		c.JSON(http.StatusCreated, "ok")
	}

}

func Logout(authClient authv1.AuthenticationServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionID, exists := c.Get(config.CookieSessionIDString)
		if !exists {
			c.AbortWithStatus(http.StatusBadRequest)
			return					
		}

		sessID, ok := sessionID.(string)
		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return					
		}

		res, err := authClient.LogoutUser(c.Request.Context(), &authv1.LogoutUserRequest{
			SessionId: sessID,
		})

		if err != nil || !res.LoggedOut {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, "logged out")
	}
}