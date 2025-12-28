package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: Unable to fund .env file")
	}

	MongoDb := os.Getenv("MONGO_URI")

	if MongoDb == "" {
		log.Fatal("MONGO_URI not set!")
	}

	fmt.Println("MongoDB URI: ", MongoDb)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return nil
	}

	// Verify connection with Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Error pinging MongoDB:", err)
		return nil
	}

	fmt.Println("Successfully connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to fund .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE_NAME: ", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}

	return collection
}
