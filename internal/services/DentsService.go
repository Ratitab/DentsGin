package services

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/golanggin/initial/shadow/internal/models/Dents"
	"gitlab.com/golanggin/initial/shadow/pkg/database/db_drivers/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "google.golang.org/protobuf/internal/detrand"
)

type DentsService struct {
	MongoDB *mongodb.MongoDB // MySQL connection instance
}

func NewDentsService(mongodb *mongodb.MongoDB) *DentsService {
	return &DentsService{
		MongoDB: mongodb,
	}
}

func (s *DentsService) GetDents() ([]Dents.Dent, error) {
	var dents []Dents.Dent

	collection := s.MongoDB.Client.Database("ratitabidze").Collection("teethImplementation")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var dent Dents.Dent

		if err := cursor.Decode(&dent); err != nil {
			return nil, err
		}
		dents = append(dents, dent)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dents, nil
}

func (s *DentsService) StorePacientsData(data Dents.PacientData) (string, error) {
	fmt.Println("akaavaaartt")
	collection := s.MongoDB.Client.Database("ratitabidze").Collection("pacients")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var existingData Dents.PacientData

	filter := bson.M{"email": data.Email}
	err := collection.FindOne(ctx, filter).Decode(&existingData)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err := collection.InsertOne(ctx, data)
			if err != nil {
				return "", nil
			}
			fmt.Println("Patient data stored successfully", data)
			return "pacient data stored successfully", nil
		}
		return "", nil
	}

	existingPhaseMap := make(map[int]bool)
	for _, phase := range existingData.Phases {
		existingPhaseMap[int(phase.ID)] = true
	}

	var newPhases []Dents.Phase
	for _, phase := range data.Phases {
		if !existingPhaseMap[int(phase.ID)] {
			newPhases = append(newPhases, phase)
		}
	}

	if len(newPhases) == 0 {
		return "no new phases to add", nil
	}

	update := bson.M{
		"$push": bson.M{"phases": bson.M{"$each": newPhases}},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", nil
	}

	fmt.Println("Patient data updated successfully with new phases: ", newPhases)
	return "Patient data updated successfully with new phases", nil
	// if err != nil {
	// 	return "", err
	// }
	// fmt.Println("patients data is stored successuflly: ", data)
	// return "Pacients Data stored successfully", nil
}
