package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//noinspection ALL
var SCHEMA = "pekabeta"
var DSN = "mongodb://username:secret@database:27017"

func InitDb() *mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI(DSN)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(SCHEMA)
}
