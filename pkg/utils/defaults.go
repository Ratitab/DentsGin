package utils

import (
	"github.com/gin-gonic/gin"
	"os"
)

func DefaultPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	return port
}

func DefaultEnvironment() string {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "local":
		gin.SetMode(gin.DebugMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	return appEnv
}
