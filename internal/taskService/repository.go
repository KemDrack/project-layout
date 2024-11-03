package taskService

import (
    
    "gorm.io/gorm"
	
    
)

// Как уже говорилось выше, репозиторий это сущность для работы с базой данных.
// Также для организации “логики” нашего приложения мы будем использовать паттерн Repository. Его суть заключается
//в том, что мы выносим всю работу с базой данных в отдельную сущность под названием репозиторий


//Он описывает что можно сделать, но не как это сделать.
type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(task Task, id uint) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}
//Этот интерфейс говорит: "любая структура, которая реализует эти методы, 
//может быть использована как TaskRepository"





type taskRepository struct { // структура для работы с БД
	dB *gorm.DB // Будет храниться подключение к бд, т.е через нее можем напрямую работать с бд
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{dB: db}
}



func(r *taskRepository) CreateTask(task Task) (Task, error) {
	result:= r.dB.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil

}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.dB.Find(&tasks).Error
	return tasks, err
}

func(r *taskRepository) UpdateTaskByID(taskNew Task, id uint) (Task, error) {
	var task Task
	if err:= r.dB.First(&task, id).Error; err!= nil {
		return Task{}, err
	}

	if err:= r.dB.Model(&task).Updates(taskNew).Error; err!= nil {
		return Task{}, err // Если обновление не удалось, возвращаем ошибку
	}

	return task, nil

}

func(r *taskRepository) DeleteTaskByID(id uint) (error) {
	if err:= r.dB.Delete(&Task{}, id).Error; err!= nil {
		return err
	}

	return nil
}



