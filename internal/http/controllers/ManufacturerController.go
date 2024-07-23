package controllers

// import (
// 	"github.com/gin-gonic/gin"
// 	"gitlab.com/golanggin/initial/shadow/internal/handlers"
// 	"gitlab.com/golanggin/initial/shadow/internal/services"
// 	"net/http"
// )

// type Controller struct {
// 	ManufacturerService *services.ManufacturerService
// }

// func ManufacturerController(manufacturerService *services.ManufacturerService) *Controller {
// 	return &Controller{
// 		ManufacturerService: manufacturerService,
// 	}
// }

// func (c *Controller) GetManufacturersHandler(ctx *gin.Context) {
// 	manufacturers, err := c.ManufacturerService.GetManufacturers()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch manufacturers"})
// 		return
// 	}

// 	handlers.GenerateResponse(ctx, manufacturers, "success", http.StatusOK)
// }
