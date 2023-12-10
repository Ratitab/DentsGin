package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/golanggin/initial/shadow/internal/http/routes"
	"gitlab.com/golanggin/initial/shadow/pkg/utils"
)

func main() {
	utils.LoadEnv()
	utils.DefaultEnvironment()

	fmt.Println("Starting the server...", gin.Mode())
	routes.SetupRoutes()
}
