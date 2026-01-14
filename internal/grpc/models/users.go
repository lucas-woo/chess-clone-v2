package models

import (
	"github.com/google/uuid"
)

var (
	DatabaseName string = "user_database"
	UserLoginCollection string = "user_login";
	UserProfileCollection string = "user_profile"
)

type UserLogin struct {
	Username string `bson:"username,omitempty"`
	Email string `bson:"email,omitempty"`
	Hash string `bson:"hash,omitempty"`
	UserID uuid.UUID `bson:"uuid,omitempty"`
}

type UserProfile struct {
	UserID uuid.UUID `bson:"uuid,omitempty"`
	//highscore
	//...profile pic
}