package repositories

import (
	"github.com/lucas-woo/chess-clone-v2/internal/grpc/models"
	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


func GetPuzzleCollection() *mongo.Collection {
	return mongodb.MongoClient.Database(models.PuzzleDatabaseName).Collection(models.PuzzleCollectionName)
}