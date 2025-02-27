package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)


var dbClient *mongo.Client

func ConnectDB() (*mongo.Client, error) {

	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	var mongoURI = os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
		return nil, fmt.Errorf("MONGO_URI is not set in .env file")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
		return nil, fmt.Errorf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
		return nil, fmt.Errorf("Error pinging MongoDB: %v", err)
	}

	dbClient = client

	log.Println("Connected to MongoDB")
	return client, nil

}


func GetClient() *mongo.Client {
    if dbClient == nil {
        log.Fatal("Database client not initialized. Call InitDB first.")
    }
    return dbClient
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("keploy_dashboard").Collection(collectionName)
}
