package models

import (
	"context"

	"github.com/lucas-woo/chess-clone-v2/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)


type User struct {
	ID bson.ObjectID
	Username string
	Hash string
	Email string
}

func Something() bson.ObjectID {
	client := mongodb.MongoClient;
	coll := client.Database("sample_restaurants").Collection("restaurants")
	newR := User{}
	ctx := context.Background()
	something, err := coll.InsertOne(ctx, newR);
	if err != nil {
		return bson.NilObjectID
	}
	return something.InsertedID.(bson.ObjectID)
}