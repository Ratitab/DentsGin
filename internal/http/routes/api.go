package routes

import (
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/http/controllers"
	"gitlab.com/golanggin/initial/shadow/internal/http/middleware"
	"gitlab.com/golanggin/initial/shadow/pkg/utils"
)

func SetupRoutes(dentsController *controllers.DentsController) {
	router := gin.Default()
	utils.LoadEnv()

	// Define routes here using router.GET(), router.POST(), etc.
	// Example:
	// router.GET("/users", handlers.GetUsers)
	memoryStore := persist.NewMemoryStore(5 * time.Minute)
	// router.GET("/api/dents", manufacturerController.GetManufacturersHandler)
	router.GET("/api/dents", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.GetDentsHandler)

	// router.GET("/api/app-version", manufacturerController.GetManufacturersHandler)
	router.POST("/api/app-login", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.LoginHandler)
	// router.GET("/api/implants", manufacturerController.GetManufacturersHandler)
	router.POST("/api/store-data", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.StoreDataHandler)
	router.GET("/api/fetch-send-data/:email", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.FetchPacientsDataHandler)
	router.GET("api/search-treatments", dentsController.SearchTreatmentsHandler)
	router.GET("api/search-diseases", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.SearchDiseasesHandler)
	router.GET("api/check-payment-status", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.CheckPaymentStatusHandler)
	router.GET("api/check-version", cache.CacheByRequestURI(memoryStore, 30*time.Second), middleware.RateLimitMiddleware(), dentsController.CheckVersionHandler)
	// router.GET("/api/order-list", manufacturerController.GetManufacturersHandler)
	// router.GET("/api/single-order/1", manufacturerController.GetManufacturersHandler)
	// router.POST("/api/is-paid", manufacturerController.GetManufacturersHandler)

	// Start listening
	router.Run(":" + utils.DefaultPort())
}
