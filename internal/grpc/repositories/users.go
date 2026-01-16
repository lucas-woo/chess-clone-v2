package repositories

import (
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetUserAuthCollection() (*mongo.Collection) {
	return mongodb.MongoClient.Database(models.UserDatabaseName).Collection(models.UserLoginCollection)
}

//needs indexing to uuid
func GetUserProfileCollection() (*mongo.Collection) {
	return mongodb.MongoClient.Database(models.UserDatabaseName).Collection(models.UserProfileCollection)
}