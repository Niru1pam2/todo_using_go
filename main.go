package main

import (
	"todo_app/db"
	"todo_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	db.InitDB()

	server.Run(":3000")
}
