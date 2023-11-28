package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MI MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (m *MongoInstance) Collection(colName string) *mongo.Collection {
	return m.DB.Collection(colName)
}

func Start(host, dbName string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", host)))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(dbName),
	}
}
