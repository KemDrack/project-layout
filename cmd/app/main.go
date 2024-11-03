package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"petproject/internal/database"
	"petproject/internal/handlers"
	"petproject/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB) // подключение к БД
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.PutTasksHandler).Methods("PUT")
	router.HandleFunc("/api/messages/{id}", handler.DeleteTaskHandler).Methods("DELETE")


	http.ListenAndServe(":8080", router)
}