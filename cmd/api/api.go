package main

import (
	"log"
	"todo_app/routes"

	"github.com/gin-gonic/gin"
)

type config struct {
	addr string
}

type application struct {
	config      config
	taskHandler *routes.TaskHandler
}

func (app *application) mount() *gin.Engine {
	router := gin.Default()

	tasksGroup := router.Group("/tasks")
	{
		// 3. Use the app.taskHandler and the Capitalized method names!
		tasksGroup.GET("/", app.taskHandler.GetTasks)
		tasksGroup.GET("/:id", app.taskHandler.GetSingleTask)
		tasksGroup.POST("/", app.taskHandler.InsertTask)
		tasksGroup.PATCH("/:id", app.taskHandler.UpdateTask)
		tasksGroup.DELETE("/:id", app.taskHandler.DeleteTask)
	}

	return router
}

func (app *application) run() {
	router := app.mount()

	log.Printf("Starting server on %s", app.config.addr)
	log.Fatal(router.Run(app.config.addr))
}
