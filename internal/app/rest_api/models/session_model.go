package models

type UserSession struct {
	SessionID string `json:"session_id" binding:"required"`
}