package handlers

import (
	"fmt"
	"net/http"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

// Глобальная переменная
var tasks = []models.Task{
	{ID: "1", Title: "Название1", Completed: false},
	{ID: "2", Title: "Название2", Completed: false},
}

// GetTasks возвращает все задачи
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// PostTask добавляет новую задачу
func PostTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	newTask.ID = fmt.Sprintf("%d", len(tasks)+1) // Простейшая генерация ID
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

// DeleteTask удаляет задачу по ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "задача удалена"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
}

// PutTask обновляет задачу по ID
func PutTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных",
		})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			ctx.JSON(http.StatusOK, updatedTask)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
}
