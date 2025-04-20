package repository

import "github/ArikuWoW/task_service/internal/models"

type Repository struct {
	TaskRepository models.TaskRepository
}

func NewRepository() *Repository {
	return &Repository{
		TaskRepository: NewInMemoryTaskRepo(),
	}
}
