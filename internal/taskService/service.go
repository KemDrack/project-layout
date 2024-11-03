// Фильтрация / проверка / работа с данными перед тем, как передать их в репозиторий
package taskService


type TaskService struct {
	rep TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{rep: repo}
}


func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.rep.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.rep.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(task Task, id uint) (Task, error) {
	return s.rep.UpdateTaskByID(task, id)
}

func (s *TaskService) DeleteTaskByID(id uint) error{
	return s.rep.DeleteTaskByID(id)
}


// func (s *TaskService) GetAllTasks() ([]Task, error) {
// 	return s.repo.GetAllTasks(task)
// }