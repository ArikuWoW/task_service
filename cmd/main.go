package main

import (
	"context"
	"github/ArikuWoW/task_service/internal/app"
	"github/ArikuWoW/task_service/internal/handler"
	"github/ArikuWoW/task_service/internal/models"
	"github/ArikuWoW/task_service/internal/repository"
	"github/ArikuWoW/task_service/pkg/logger"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger.InitLogger(true)
	rand.Seed(time.Now().UnixNano())
	defer logger.Log.Sync()

	logger.Log.Info("Logger initialized")

	repo := repository.NewInMemoryTaskRepo()

	service := app.NewService(repo)

	pool := repository.NewWorkerPool(service.Task, 100)
	pool.Start(4)

	handler := handler.NewHandler(service, pool)

	srv := new(models.Server)
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			logger.Log.Fatal("Failed to run HTTP server", zap.Error(err))
		}
	}()

	logger.Log.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Log.Info("App shutting sown")

	pool.Wait()

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Log.Fatal("Error server shutdown", zap.Error(err))
	}
}
