package main

import (
	"log"
	"net/http"
	"petproject/internal/database"
	"petproject/internal/handlers"
	"petproject/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	// database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB) // соединение с базой данных
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.PutTasksHandler).Methods("PUT")
	router.HandleFunc("/api/messages/{id}", handler.DeleteTaskHandler).Methods("DELETE")


	if err := http.ListenAndServe(":8080", router);err!= nil {
		log.Printf("Ошибка в запуске сервера %v" , err)
	}
}