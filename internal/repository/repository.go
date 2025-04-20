package repository

import "github/ArikuWoW/task_service/internal/models"

type TaskRepository interface {
	Save(task *models.Task) error
	FindByID(id string) (*models.Task, error)
	Update(task *models.Task) error
}

type Repository struct {
	TaskRepository
}

func NewRepository() *Repository {
	return &Repository{
		TaskRepository: NewInMemoryTaskRepo(),
	}
}
