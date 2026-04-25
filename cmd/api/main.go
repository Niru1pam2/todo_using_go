package main

import (
	"log"
	"todo_app/internal/db"
	"todo_app/internal/store"
	"todo_app/routes"
	"todo_app/service"
)

func main() {
	dbConn, err := db.InitDB("api.db")

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	defer dbConn.Close()

	storage := store.NewStorage(dbConn)
	taskService := service.NewTaskService(storage)

	myTaskHandler := routes.NewTaskHandler(taskService)

	app := &application{
		config: config{
			addr: ":3000",
		},
		taskHandler: myTaskHandler,
	}

	app.run()

}
