package mongodb

import (
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client;

func ConnectMongo() (err error) {
	uri := os.Getenv("MONGO_URI");
	if uri == "" {
		err = errors.New("enable to connect to mongo")
		return 
	}

	options := options.Client().ApplyURI(uri);
	MongoClient, err = mongo.Connect(options)

	return 

}