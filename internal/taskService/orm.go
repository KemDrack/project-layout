package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model // добавляем поля ID, CreatedAt, UpdatedAt и DeletedAt
	Task string `json:"task"`  // Вот эти поля task и progress будут отображаться в бд
	IsDone bool `json:"is_done"`
}

//  структура нашего Message для БД, то, какие столбцы будут в нашей БД