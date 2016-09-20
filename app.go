package main

import (
	"github.com/gin-gonic/gin"
	"gotodolist/models"
	"log"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/tasks", GetTasks)
		v1.POST("/lists/:listId/tasks", PostTask)
	}
	log.SetFlags(log.LstdFlags)
	r.Run(":8080")
}

func GetTasks(context *gin.Context) {

	task1 := models.Task{"My first task", "My simple description"}

	list := models.TodoList{}
	list.Name = "My list"

	list.AddTask(task1)
	context.JSON(200, list)
}

func PostTask(context *gin.Context) {

	task := models.Task{}
	context.Bind(&task)
	//listId := context.Param("listId")

	//log.Println("This is the list id: " + listId)

	context.JSON(200, task)
}
