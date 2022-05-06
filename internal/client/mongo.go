package client

import (
	"context"
	"getir-assignment/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(config config.MongoConfig) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.URI))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	return client
}
