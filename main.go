package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func main() {
	isRunningOnVercel := os.Getenv("VERCEL") == "1"
	if isRunningOnVercel {
		fmt.Println("🚀 Running on Vercel (serverless mode)...")
		return
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Todo{})

	apiRouter := gin.Default()

	routes.SetupTodoRoutes(apiRouter)

	serverPort := os.Getenv("PORT")

	fmt.Println("🚀 Server is running locally on port", serverPort)
	if err := apiRouter.Run(":" + serverPort); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
