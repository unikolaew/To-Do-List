package main

import (
	"todo-list/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.PostTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)
	r.PUT("/tasks/:id", handlers.PutTask)

	r.Run(":8080")
}
