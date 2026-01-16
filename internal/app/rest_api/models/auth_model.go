package models

type UserLogin struct {

}

type UserSignUpData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	RememberMe bool `json:"rememberMe" binding:"required"`
}