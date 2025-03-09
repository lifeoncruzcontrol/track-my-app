package db

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
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

// loadEnv manually reads a .env file and sets environment variables.
func loadEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines and comments
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Invalid line in .env file: %s", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		value = strings.Trim(value, `"'`)

		// Set environment variable
		os.Setenv(key, value)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading .env file: %v", err)
	}
}

func init() {

	loadEnv(".env")
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
