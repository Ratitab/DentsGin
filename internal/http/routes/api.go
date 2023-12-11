package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/http/controllers"
	"gitlab.com/golanggin/initial/shadow/pkg/utils"
)

func SetupRoutes(manufacturerController *controllers.Controller) {
	router := gin.Default()
	utils.LoadEnv()

	// Define routes here using router.GET(), router.POST(), etc.
	// Example:
	// router.GET("/users", handlers.GetUsers)

	router.GET("/api/manufacturer-filters", manufacturerController.GetManufacturersHandler)

	// Start listening
	router.Run(":" + utils.DefaultPort())
}
