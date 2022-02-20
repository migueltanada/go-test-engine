package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// for mongodb
var db *mongo.Database

func GetMongoDB(ctx context.Context, host string, dbName string) *mongo.Database {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		os.Exit(1)
	}

	db := client.Database("test")
	return db
}
