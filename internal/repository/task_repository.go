package repository

import (
	"errors"
	"github/ArikuWoW/task_service/internal/models"
	"sync"
)

type InMemoryTaskRepo struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{
		tasks: make(map[string]*models.Task),
	}
}

func (r *InMemoryTaskRepo) Save(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepo) FindByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepo) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[task.ID]
	if !ok {
		return errors.New("task not found")
	}

	r.tasks[task.ID] = task
	return nil
}
