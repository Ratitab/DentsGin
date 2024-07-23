package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/http/controllers"
	"gitlab.com/golanggin/initial/shadow/pkg/utils"
)

func SetupRoutes(dentsController *controllers.DentsController) {
	router := gin.Default()
	utils.LoadEnv()

	// Define routes here using router.GET(), router.POST(), etc.
	// Example:
	// router.GET("/users", handlers.GetUsers)

	// router.GET("/api/dents", manufacturerController.GetManufacturersHandler)
	router.GET("/api/dents", dentsController.GetDentsHandler)

	// router.GET("/api/app-version", manufacturerController.GetManufacturersHandler)
	router.POST("/api/app-login", dentsController.LoginHandler)
	// router.GET("/api/implants", manufacturerController.GetManufacturersHandler)
	router.POST("/api/store-data", dentsController.StoreDataHandler)
	router.GET("/api/fetch-send-data/:email", dentsController.FetchPacientsDataHandler)
	router.GET("api/search-treatments", dentsController.SearchTreatmentsHandler)
	router.GET("api/search-diseases", dentsController.SearchDiseasesHandler)
	router.GET("api/check-payment-status", dentsController.CheckPaymentStatusHandler)
	router.GET("api/check-version", dentsController.CheckVersionHandler)
	// router.GET("/api/order-list", manufacturerController.GetManufacturersHandler)
	// router.GET("/api/single-order/1", manufacturerController.GetManufacturersHandler)
	// router.POST("/api/is-paid", manufacturerController.GetManufacturersHandler)

	// Start listening
	router.Run(":" + utils.DefaultPort())
}
