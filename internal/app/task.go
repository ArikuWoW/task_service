package app

import (
	"github/ArikuWoW/task_service/internal/models"
	"time"

	"github.com/google/uuid"
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
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetTask(id string) (*models.Task, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) UpdateTaskResult(id string, status models.Status, result string, errMsg string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	task.Status = status
	task.Result = result
	task.Error = errMsg
	task.UpdatedAt = time.Now()

	return s.repo.Update(task)
}
