package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.GET("/:id", controllers.GetTodoByID)
		todoRoutes.PUT("/:id", controllers.UpdateTodoByID)
		todoRoutes.DELETE("/:id", controllers.DeleteTodoByID)
	}
}
