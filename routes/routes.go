package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", getTasks)
	server.GET("/:id", getSingleTask)
	server.POST("/", insertTask)
	server.PATCH("/:id", updateTask)
	server.DELETE("/:id", deleteTask)
}
