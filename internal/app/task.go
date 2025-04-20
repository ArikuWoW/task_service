package app

import (
	"github/ArikuWoW/task_service/internal/models"
	"github/ArikuWoW/task_service/pkg/logger"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ models.TaskProcessor = (*TaskService)(nil)

type TaskService struct {
	repo models.TaskRepository
}

func NewTaskService(repo models.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(taskType string) (*models.Task, error) {
	now := time.Now()
	task := &models.Task{
		ID:        uuid.NewString(),
		Type:      taskType,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    models.StatusWait,
	}

	if err := s.repo.Save(task); err != nil {
		logger.Log.Error("Failed to save task",
			zap.String("task_id", task.ID),
			zap.Error(err),
		)

		return nil, err
	}

	logger.Log.Info("Task created",
		zap.String("task_id", task.ID),
		zap.String("type", task.Type),
	)

	return task, nil
}

func (s *TaskService) GetTask(id string) (*models.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		logger.Log.Warn("Task not found",
			zap.String("task_id", id),
			zap.Error(err),
		)

		return nil, err
	}

	logger.Log.Debug("Task fetched",
		zap.String("task_id", task.ID),
		zap.String("status", string(task.Status)),
	)

	return task, nil
}

func (s *TaskService) UpdateTaskResult(id string, status models.Status, result string, errMsg string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		logger.Log.Error("Failed to find task for update",
			zap.String("task_id", id),
			zap.Error(err),
		)
		return err
	}

	task.Status = status
	task.Result = result
	task.Error = errMsg
	task.UpdatedAt = time.Now()

	if err := s.repo.Update(task); err != nil {
		logger.Log.Error("Failed to update task",
			zap.String("task_id", task.ID),
			zap.Error(err),
		)
		return err
	}

	logger.Log.Info("Task result updated",
		zap.String("task_id", task.ID),
		zap.String("status", string(status)),
		zap.String("result", result),
		zap.String("error", errMsg),
	)

	return nil
}
