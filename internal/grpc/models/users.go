package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/google/uuid"
)

var (
	DatabaseName string = "user_database"
	UserLoginCollection string = "user_login";
	UserProfileCollection string = "user_profile"
)

type UserLogin struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
	Email string `bson:"email,omitempty"`
	Hash string `bson:"hash,omitempty"`
	UserID uuid.UUID `bson:"uuid,omitempty"`
}

type UserProfile struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	UserID uuid.UUID `bson:"uuid,omitempty"`
	//highscore
}