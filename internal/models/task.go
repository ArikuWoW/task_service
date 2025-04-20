package models

import "time"

type Status string

const (
	StatusWait Status = "WAITING"
	StatusRun  Status = "RUNNING"
	StatusDone Status = "DONE"
	StatusFail Status = "FAILED"
)

type Task struct {
	ID        string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status
	Result    string
	Error     string
}

type TaskRepository interface {
	Save(task *Task) error
	FindByID(id string) (*Task, error)
	Update(task *Task) error
}

type TaskProcessor interface {
	UpdateTaskResult(id string, status Status, result string, errMsg string) error
}
