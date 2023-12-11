package services

import (
	"gitlab.com/golanggin/initial/shadow/pkg/database/db_drivers/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

import (
	"context"
)

type UserLogsService struct {
	MongoDB *mongodb.MongoDB // MongoDB connection instance
}

func NewUserLogsService(mongoDB *mongodb.MongoDB) *UserLogsService {
	return &UserLogsService{
		MongoDB: mongoDB,
	}
}

type YourLogStruct struct {
	ID            string `bson:"_id"`
	UserID        string `bson:"user_id"`
	Logable       string `bson:"logable"`
	LogableID     string `bson:"logable_id"`
	TextLog       string `bson:"text_log"`
	UserLogTypeID int    `bson:"user_log_type_id"`
	IsSystem      bool   `bson:"is_system"`
}

func (s *UserLogsService) GetUserLogs() ([]YourLogStruct, error) {
	// Assuming YourLogStruct is your structure for log data
	collection := s.MongoDB.Client.Database("hrmongo").Collection("user_logs")

	// Query the MongoDB collection and fetch user logs
	ctx := context.TODO()                         // Create a context
	cursor, err := collection.Find(ctx, bson.M{}) // Fetch all documents in the collection
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []YourLogStruct

	// Iterate through the cursor and extract log data into YourLogStruct
	for cursor.Next(ctx) {
		var log YourLogStruct
		if err := cursor.Decode(&log); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return logs, nil // Return logs fetched from the collection
}
