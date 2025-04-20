package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	taskType := c.Query("type")
	if taskType == "" {
		taskType = "default"
	}

	t, err := h.service.Task.CreateTask(taskType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.pool.AddTaskToStack(t)

	c.JSON(http.StatusAccepted, t)
}

func (h *Handler) getTask(c *gin.Context) {
	id := c.Param("id")
	t, err := h.service.Task.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, t)
}
