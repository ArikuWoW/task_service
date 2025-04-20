package handler

import (
	"github/ArikuWoW/task_service/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) createTask(c *gin.Context) {
	taskType := c.Query("type")
	if taskType == "" {
		taskType = "default"
	}

	t, err := h.service.Task.CreateTask(taskType)
	if err != nil {
		logger.Log.Error("Failed to create task",
			zap.String("type", taskType),
			zap.Error(err),
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Log.Info("Task created",
		zap.String("task_id", t.ID),
		zap.String("type", t.Type),
	)

	h.pool.AddTaskToStack(t)

	c.JSON(http.StatusAccepted, t)
}

func (h *Handler) getTask(c *gin.Context) {
	id := c.Param("id")
	t, err := h.service.Task.GetTask(id)
	if err != nil {
		logger.Log.Warn("Task not found",
			zap.String("task_id", id),
			zap.Error(err),
		)

		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	logger.Log.Info("Task sended",
		zap.String("task_id", t.ID),
		zap.String("status", string(t.Status)),
	)

	c.JSON(http.StatusOK, t)
}
