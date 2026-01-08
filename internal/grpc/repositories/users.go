package repositories

import (
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetUserLoginCollection() (*mongo.Collection) {
	return mongodb.MongoClient.Database(models.DatabaseName).Collection(models.UserLoginCollection)
}

func GetUserProfileCollection() (*mongo.Collection) {
	return mongodb.MongoClient.Database(models.DatabaseName).Collection(models.UserProfileCollection)
}