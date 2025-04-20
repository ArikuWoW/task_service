package repository

import (
	"github/ArikuWoW/task_service/internal/models"
	"log"
	"sync"
	"time"
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
		log.Printf("[Worker %d] Processing task: %s", id, task.ID)

		time.Sleep(30 * time.Second)

		err := wp.processor.UpdateTaskResult(task.ID, models.StatusDone, "result: OK", "")
		if err != nil {
			log.Printf("[Worker %d] Failed task %s: %v", id, task.ID, err)
		} else {
			log.Printf("[Worker %d] Complete task %s: %v", id, task.ID)
		}

		wp.wg.Done()
	}
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
