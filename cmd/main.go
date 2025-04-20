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
	// Инициализация логгер
	logger.InitLogger(true)
	defer logger.Log.Sync()

	rand.Seed(time.Now().UnixNano())

	logger.Log.Info("Logger initialized")

	// Создание репозитория в памяти (в дальнейшем можно поменять на БД)
	repo := repository.NewInMemoryTaskRepo()

	// Инициализация сервисов
	service := app.NewService(repo)

	// Создание и запуск воркер пула и 5 воркерами
	pool := repository.NewWorkerPool(service.Task, 100)
	pool.Start(5)

	// Создание хендлера и подклбчения сервиса с воркер пулом
	handler := handler.NewHandler(service, pool)

	// Запуск сервера асинхронным методом что бы  в дальнейшем использовать graceful shutdown
	srv := new(models.Server)
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			logger.Log.Fatal("Failed to run HTTP server", zap.Error(err))
		}
	}()

	logger.Log.Info("App started")

	// Обработка ctrl+c для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Получаем из канала сигнал о завершении
	<-quit

	logger.Log.Info("App shutting sown")

	// Ждем завершения всех задач из воркер пула и завершаем сервер
	pool.Wait()

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Log.Fatal("Error server shutdown", zap.Error(err))
	}
}
