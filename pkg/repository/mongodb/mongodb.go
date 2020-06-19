package mongodb

import (
	"context"
	"starwars-hex/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Open opens a connection to the especified mongodb uri and database.
func Open(uri, database string) *mongo.Database {
	if db != nil {
		Close()
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal("mongodb.Open", "mongo.NewClient", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal("mongodb.Open", "client.Connect", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal("mongodb.Open", "client.Ping", err)
	}

	db = client.Database(database)

	createIndex()

	return db
}

// Close closes the current connection.
func Close() {
	if db != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		err := db.Client().Disconnect(ctx)
		if err != nil {
			logger.Error("mongodb.Close", "db.Client.Disconnect", err)
		}
	}
}

func createIndex() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	planetIdxOpts := options.Index().SetName("idx_planets_name").SetUnique(true)

	_, err := db.Collection("planets").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: planetIdxOpts})

	if err != nil {
		logger.Fatal("mongodb.createIndex", "db.Indexes.CreateOne", err, planetIdxOpts.Name)
	}

	return nil
}
