package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/pkg/utils"
)

func SetupRoutes() {
	router := gin.Default()
	utils.LoadEnv()

	// Define routes here using router.GET(), router.POST(), etc.
	// Example:
	// router.GET("/users", handlers.GetUsers)

	// Start listening
	router.Run(":" + utils.DefaultPort())
}
