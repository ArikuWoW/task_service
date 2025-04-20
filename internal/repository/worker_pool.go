package repository

import (
	"github/ArikuWoW/task_service/internal/models"
	"github/ArikuWoW/task_service/pkg/logger"
	"math/rand"
	"sync"
	"time"

	"go.uber.org/zap"
)

type WorkerPool struct {
	processor models.TaskProcessor
	queue     chan *models.Task
	wg        sync.WaitGroup
}

func NewWorkerPool(processor models.TaskProcessor, buff int) *WorkerPool {
	return &WorkerPool{
		processor: processor,
		queue:     make(chan *models.Task, buff),
	}
}

func (wp *WorkerPool) Start(n int) {
	for i := 0; i < n; i++ {
		go wp.worker(i)
	}
}

func (wp *WorkerPool) AddTaskToStack(task *models.Task) {
	wp.wg.Add(1)
	wp.queue <- task
}

func (wp *WorkerPool) worker(id int) {
	for task := range wp.queue {
		logger.Log.Info("Worker processing task",
			zap.Int("worker_id", id),
			zap.String("task_id", task.ID),
		)

		// Имитация времени работы реального приложения
		// Получаем Ответ от 2-5 минут рандомно
		delay := time.Duration(2+rand.Intn(4)) * time.Second

		logger.Log.Info("Worker sleeping",
			zap.Int("worker_id", id),
			zap.String("task_id", task.ID),
			zap.Duration("sleep_duration", delay),
		)

		time.Sleep(delay)

		err := wp.processor.UpdateTaskResult(task.ID, models.StatusDone, "result: OK", "")
		if err != nil {
			logger.Log.Error("Failed to update task result",
				zap.Int("worker_id", id),
				zap.String("task_id", task.ID),
				zap.Error(err),
			)
		} else {
			logger.Log.Info("Task completed successfully",
				zap.Int("worker_id", id),
				zap.String("task_id", task.ID),
			)
		}

		wp.wg.Done()
	}
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
