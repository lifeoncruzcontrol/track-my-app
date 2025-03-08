package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURI string
	Ctx      context.Context
	Cancel   context.CancelFunc
	Client   *mongo.Client
	Job_Apps *mongo.Collection
)

func init() {
	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		log.Fatal("Missing MONGO_URI variable")
		return
	}

	Ctx, Cancel = context.WithTimeout(context.Background(), 20*time.Second)

	var err error
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(MongoURI).SetMaxPoolSize(50))
	if err != nil {
		log.Fatal("Error connecting to database")
		Cancel()
		return
	}

	db := Client.Database("job-app-db")
	Job_Apps = db.Collection("job-apps")
}
