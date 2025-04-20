package app

import (
	"github/ArikuWoW/task_service/internal/models"
)

// Логика над тасками
type Task interface {
	CreateTask(taskType string) (*models.Task, error)
	GetTask(id string) (*models.Task, error)
	UpdateTaskResult(id string, status models.Status, result string, errMsg string) error
}

// Агрегатор, который в дальнейшем можно увеличивать добавляя новую логику
type Service struct {
	Task
}

func NewService(repos models.TaskRepository) *Service {
	return &Service{
		Task: NewTaskService(repos),
	}
}
