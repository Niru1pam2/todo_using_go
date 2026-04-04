package routes

import (
	"net/http"
	"strconv"
	"todo_app/models"

	"github.com/gin-gonic/gin"
)

func getTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Could not fetch tasks",
			"error":    err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func getSingleTask(context *gin.Context) {
	var task models.Task
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not parse event id.",
			"error":    err.Error(),
		})
		return
	}

	task.Id = taskId
	err = task.GetTask()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not find Task.",
			"error":    err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"Task": task,
	})
}

func insertTask(context *gin.Context) {
	var task models.Task

	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
		return
	}

	task.SaveTask()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Added task to Task list",
	})
}

func updateTask(context *gin.Context) {
	var task models.Task
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not parse event id.",
			"error":    err.Error(),
		})
		return
	}

	task.Id = taskId

	err = context.ShouldBindJSON(&task)

	task.UpdateTask()

	context.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})

}

func deleteTask(context *gin.Context) {
	var task models.Task
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not parse event id.",
			"error":    err.Error(),
		})
		return
	}

	task.Id = taskId

	err = task.DeleteTask()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not delete event.",
			"error":    err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})

}
