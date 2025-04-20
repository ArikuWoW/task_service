package handler

import (
	"github/ArikuWoW/task_service/internal/app"
	"github/ArikuWoW/task_service/internal/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *app.Service
	pool    *repository.WorkerPool
}

func NewHandler(service *app.Service, pool *repository.WorkerPool) *Handler {
	return &Handler{service: service, pool: pool}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	task := router.Group("/task")
	{
		task.POST("/create", h.createTask)
		task.GET("/:id", h.getTask)
	}

	return router
}
