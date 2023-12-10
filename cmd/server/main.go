package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/http/routes"
	"gitlab.com/golanggin/initial/shadow/pkg/database/db_drivers/mongodb"
	"gitlab.com/golanggin/initial/shadow/pkg/database/db_drivers/mysql"
	"gitlab.com/golanggin/initial/shadow/pkg/utils"
)

func main() {
	utils.LoadEnv()
	utils.DefaultEnvironment()

	fmt.Println("Starting the server... mode:", gin.Mode())
	// Connect to MySQL
	mysql := &mysql.MySQL{}
	err := mysql.Connect()
	if err != nil {
		fmt.Println("Error connecting to MySQL:", err)
	}

	// Connect to MongoDB
	mongo := &mongodb.MongoDB{}
	err = mongo.Connect()
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}
	routes.SetupRoutes()
}
