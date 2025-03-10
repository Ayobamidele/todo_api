package handler

import (
	"fmt"
	"net/http"
	"todo-api/config"
	"todo-api/models"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
)

func HandleRequest(responseWriter http.ResponseWriter, request *http.Request) {
	config.ConnectDatabase()

	apiRouter := gin.Default()

	routes.SetupTodoRoutes(apiRouter)

	apiRouter.ServeHTTP(responseWriter, request)
}

func init() {
	fmt.Println("🚀 Initializing application...")
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Todo{})
}
