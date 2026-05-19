package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/google/uuid"
)

var (
	UserDatabaseName string = "user_database"
	UserLoginCollection string = "user_login";
	UserProfileCollection string = "user_profile"
	UserRoleCollection string = "user_role"
)

type UserLogin struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	Email string `bson:"email,omitempty"`
	Hash string `bson:"hash,omitempty"`
	UserID uuid.UUID `bson:"uuid,omitempty"` //SAME
}

type UserRole struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	UserID uuid.UUID `bson:"uuid,omitempty"` //SAME
	Role string `bson:"role,omitempty"`
}

type UserProfile struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	UserID uuid.UUID `bson:"uuid,omitempty"` // SAME
	Username string `bson:"username,omitempty"`
	HighScore uint32 `bson:"highScore,omitempty"`
	//highscore
}