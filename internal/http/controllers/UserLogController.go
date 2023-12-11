package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/services"
	"net/http"
)

type UserLogsController struct {
	UserLogsService *services.UserLogsService
}

func NewUserLogsController(userLogsService *services.UserLogsService) *UserLogsController {
	return &UserLogsController{
		UserLogsService: userLogsService,
	}
}

func (c *UserLogsController) GetUserLogsHandler(ctx *gin.Context) {
	logs, err := c.UserLogsService.GetUserLogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user logs"})
		return
	}

	ctx.JSON(http.StatusOK, logs)
}
