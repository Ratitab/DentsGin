package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	// Import other MongoDB packages as needed
)

type MongoDB struct {
	Client *mongo.Client
	// Add other MongoDB-specific fields here
}

func (m *MongoDB) Connect() error {
	connectionURI := os.Getenv("MONGO_SRV")
	if connectionURI == "" {
		// MongoDB connection parameters
		dbHost := os.Getenv("MONGO_DB_HOST")
		dbPort := os.Getenv("MONGO_DB_PORT")
		dbName := os.Getenv("MONGO_DB_COLLECTION")
		dbUser := os.Getenv("MONGO_DB_USERNAME")
		dbPass := os.Getenv("MONGO_DB_PASSWORD")

		connectionURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	clientOptions := options.Client().ApplyURI(connectionURI)

	// Create a context with timeout
	ctx := context.TODO()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the MongoDB server
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// Assign the connected client to the struct field
	m.Client = client

	fmt.Println("Connected to MongoDB Atlas!")
	return nil
}
