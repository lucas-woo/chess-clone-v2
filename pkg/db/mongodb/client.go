package mongodb

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var MongoClient *mongo.Client;

func ConnectMongo() (err error) {
	
	initContext, timoutFunc := context.WithTimeout(context.Background(), time.Second * 10);

	defer timoutFunc();

	uri, found := os.LookupEnv("MONGO_URI");
	if !found {
		err = errors.New("enable to connect to mongo")
		return 
	}

	options := options.Client().ApplyURI(uri);
	MongoClient, err = mongo.Connect(options);

	if err != nil {
		return
	}

	err = MongoClient.Ping(initContext, readpref.PrimaryPreferred())

	return 

}