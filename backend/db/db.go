package db

import (
	"log"
	"os"
)

var (
	MongoURI string
)

func init() {
	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		log.Fatal("Missing MONGO_URI variable")
		return
	}
}
