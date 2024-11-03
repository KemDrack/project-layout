package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"petproject/internal/taskService" // Импортируем наш сервис
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}


func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service,}
}

//Таким образом мы убрали логику работы с БД из нашей ручки, что является хорошей
// практикой и правильным подходом в целом. Теперь мы просто обращаемся к функции 
//сервиса GetAllTasks, которая обращается в репозиторий, который возвращает все задачи.
func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}


//Тут мы также передаем дешифрованный task в сервис, который в свою очередь уже 
//возвращает createdTask, который мы возвращаем в качестве ответа нашей ручки
func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}


func (h *Handler) PutTasksHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
    idStr := vars["id"]


	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Println("Invalid ID:", idStr)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }


	var updatedTask taskService.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.Service.UpdateTaskByID(updatedTask, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}


func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	w.WriteHeader(http.StatusNoContent)

}




