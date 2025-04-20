package app

import (
	"github/ArikuWoW/task_service/internal/models"
)

type Task interface {
	CreateTask(taskType string) (*models.Task, error)
	GetTask(id string) (*models.Task, error)
	UpdateTaskResult(id string, status models.Status, result string, errMsg string) error
}

type Service struct {
	Task
}

func NewService(repos models.TaskRepository) *Service {
	return &Service{
		Task: NewTaskService(repos),
	}
}
