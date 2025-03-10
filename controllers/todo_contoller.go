package controllers

import (
	"net/http"
	"todo-api/config"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(context *gin.Context) {
	var newTodo models.Todo
	if err := context.ShouldBindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newTodo)
	context.JSON(http.StatusCreated, newTodo)
}

func GetTodos(context *gin.Context) {
	var todoList []models.Todo
	config.DB.Find(&todoList)
	context.JSON(http.StatusOK, todoList)
}

func GetTodoByID(context *gin.Context) {
	todoID := context.Param("id")
	var todoItem models.Todo
	if err := config.DB.First(&todoItem, todoID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, todoItem)
}

func UpdateTodoByID(context *gin.Context) {
	todoID := context.Param("id")
	var existingTodo models.Todo
	if err := config.DB.First(&existingTodo, todoID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := context.ShouldBindJSON(&existingTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&existingTodo)
	context.JSON(http.StatusOK, existingTodo)
}

func DeleteTodoByID(context *gin.Context) {
	todoID := context.Param("id")
	if err := config.DB.Delete(&models.Todo{}, todoID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
