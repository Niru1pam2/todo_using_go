package routes

import (
	"net/http"
	"strconv"
	"todo_app/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(context *gin.Context) {
	tasks, err := h.service.GetAllTasks()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Could not fetch tasks",
			"error":    err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetSingleTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not parse event id.",
			"error":    err.Error(),
		})
		return
	}

	task, err := h.service.GetSingleTask(taskId)

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

func (h *TaskHandler) InsertTask(context *gin.Context) {

	var requestBody struct {
		Title      string `json:"title" binding:"required"`
		IsFinished bool   `json:"isFinished"`
	}

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
		return
	}

	h.service.CreateTask(requestBody.Title, requestBody.IsFinished)

	context.JSON(http.StatusCreated, gin.H{
		"message": "Added task to Task list",
	})
}

func (h *TaskHandler) UpdateTask(context *gin.Context) {
	var requestBody struct {
		Title      string `json:"title" binding:"required"`
		IsFinished bool   `json:"isFinished"`
	}

	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not parse event id.",
			"error":    err.Error(),
		})
		return
	}

	err = context.ShouldBindJSON(&requestBody)

	h.service.UpdateTask(taskId, requestBody.Title, requestBody.IsFinished)

	context.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})

}

func (h *TaskHandler) DeleteTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not parse event id.",
			"error":    err.Error(),
		})
		return
	}

	err = h.service.DeleteTask(taskId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Could not delete event.",
			"error":    err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})

}
