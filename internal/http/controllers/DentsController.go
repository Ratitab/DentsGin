package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/handlers"
	"gitlab.com/golanggin/initial/shadow/internal/models/Dents"
	"gitlab.com/golanggin/initial/shadow/internal/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DentsController struct {
	Controller                          // Embed the Controller interface
	DentsService *services.DentsService // Specific service for manufacturers
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type LoginResponse struct {
	Success bool `json:"success" bson:"success"`
}

func (mc *DentsController) SetService(service interface{}) {
	if svc, ok := service.(*services.DentsService); ok {
		mc.DentsService = svc
	} else {
		panic("Invalid service type provided for ManufacturerController")
	}
}

func (c *DentsController) GetDentsHandler(ctx *gin.Context) {
	dents, err := c.DentsService.GetDents()
	if err != nil {
		handlers.GenerateResponse(ctx, nil, "Failed to fetch manufacturers", http.StatusInternalServerError)
		return
	}
	handlers.GenerateResponse(ctx, dents, "success", http.StatusOK)
}

func (c *DentsController) LoginHandler(ctx *gin.Context) {
	var creds Credentials
	if err := ctx.BindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invaliud request payload"})
		return
	}

	collection := c.DentsService.MongoDB.Client.Database("ratitabidze").Collection("teethImplementation")

	var result Credentials
	err := collection.FindOne(context.TODO(), bson.M{"username": creds.Username, "password": creds.Password}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusUnauthorized, LoginResponse{Success: false})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	log.Println("credentilas is: ", result)
	ctx.JSON(http.StatusOK, LoginResponse{Success: true})
}

func (c *DentsController) StoreDataHandler(ctx *gin.Context) {
	var data Dents.PacientData

	if err := ctx.ShouldBindJSON(&data); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("recieved data for storage: ", data)
	fmt.Println("recieved data for storage: ", data)
	// panic("NIIICEE")
	message, err := c.DentsService.StorePacientsData(data)
	if err != nil {
		log.Println("Error storing pacient data:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Pacient data stored successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (c *DentsController) FetchPacientsDataHandler(ctx *gin.Context) {
	log.Println("FetchPacientsDataHandler started")

	email := ctx.Param("email")
	log.Println("Received email:", email)

	var data Dents.PacientData
	collection := c.DentsService.MongoDB.Client.Database("ratitabidze").Collection("pacients")

	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&data)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *DentsController) SearchTreatmentsHandler(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is requiered"})
		return
	}

	collection := c.DentsService.MongoDB.Client.Database("ratitabidze").Collection("treatments")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": bson.M{"$regex": query, "$options": "i"}}
	opts := options.Find().SetLimit(10)

	cursor, err := collection.Find(ctxMongo, filter, opts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(ctxMongo)

	var results []Dents.SearchItem

	for cursor.Next(ctxMongo) {
		var item Dents.SearchItem
		if err := cursor.Decode(&item); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, item)
	}
	fmt.Println(results)
	ctx.JSON(http.StatusOK, results)
}

func (c *DentsController) SearchDiseasesHandler(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	collection := c.DentsService.MongoDB.Client.Database("ratitabidze").Collection("diseases")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": bson.M{"$regex": query, "$options": "i"}}
	opts := options.Find().SetLimit(10)

	cursor, err := collection.Find(ctxMongo, filter, opts)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(ctxMongo)

	var results []Dents.SearchItem
	for cursor.Next(ctxMongo) {
		var item Dents.SearchItem
		if err := cursor.Decode(&item); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, item)
	}

	ctx.JSON(http.StatusOK, results)
}

func (c *DentsController) CheckPaymentStatusHandler(ctx *gin.Context) {
	collection := c.DentsService.MongoDB.Client.Database("ratitabidze").Collection("payments")

	var paymentStatus Dents.PaymentStatus

	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&paymentStatus)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No payment status document found")
			ctx.JSON(http.StatusOK, gin.H{"isPaid": false})
		} else {
			log.Println("Error finding payment status:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	fmt.Println("[ISPAID]", paymentStatus.IsPaid)
	ctx.JSON(http.StatusOK, gin.H{"isPaid": paymentStatus.IsPaid})
}

func (c *DentsController) CheckVersionHandler(ctx *gin.Context) {
	collection := c.DentsService.MongoDB.Client.Database("ratitabidze").Collection("version")

	var versionDoc struct {
		Version string `bson:"version"`
	}

	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&versionDoc)
	if err != nil {
		log.Println("Error finding version document:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	response := Dents.VersionResponse{
		Version:     versionDoc.Version,
		DownloadURL: "asfasdfws",
	}

	fmt.Println("VERSION", response.Version)
	ctx.JSON(http.StatusOK, response)
}

// func (c *DentsController) GetCitiesHandler(ctx *gin.Context) {
// 	cities, err := c.DentsSer
// 	if err != nil {
// 		handlers.GenerateResponse(ctx, nil, "Failed to fetch manufacturers", http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.GenerateResponse(ctx, cities, "success", http.StatusOK)
// }
