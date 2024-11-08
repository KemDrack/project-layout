package main

import (
	"log"
	"petproject/internal/database"
	"petproject/internal/handlers"
	"petproject/internal/taskService"

	
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB) // соединение с базой данных
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	e:= echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler:= tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err:= e.Start(":8080"); err!= nil {
		log.Fatalf("failed to start with err: %v", err)
	}
	
}