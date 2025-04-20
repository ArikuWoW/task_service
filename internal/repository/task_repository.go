package repository

import (
	"errors"
	"github/ArikuWoW/task_service/internal/models"
	"github/ArikuWoW/task_service/pkg/logger"
	"sync"

	"go.uber.org/zap"
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

	logger.Log.Info(
		"Task saved",
		zap.String("task_id", task.ID),
		zap.String("status", string(task.Status)),
	)

	return nil
}

func (r *InMemoryTaskRepo) FindByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		logger.Log.Warn("Task not found", zap.String("task_id", id))
		return nil, errors.New("task not found")
	}

	logger.Log.Debug("Task found", zap.String("task_id", id))
	return task, nil
}

func (r *InMemoryTaskRepo) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[task.ID]
	if !ok {
		logger.Log.Error("Cannot update task, not found", zap.String("task_id", task.ID))
		return errors.New("task not found")
	}

	r.tasks[task.ID] = task

	logger.Log.Info("Task updated",
		zap.String("task_id", task.ID),
		zap.String("status", string(task.Status)),
	)

	return nil
}
