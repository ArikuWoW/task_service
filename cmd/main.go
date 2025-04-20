package main

import (
	"context"
	"github/ArikuWoW/task_service/internal/app"
	"github/ArikuWoW/task_service/internal/handler"
	"github/ArikuWoW/task_service/internal/models"
	"github/ArikuWoW/task_service/internal/repository"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	repo := repository.NewInMemoryTaskRepo()

	service := app.NewService(repo)

	pool := repository.NewWorkerPool(service.Task, 100)
	pool.Start(4)

	handler := handler.NewHandler(service, pool)

	srv := new(models.Server)
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Println("App shutting sown")

	pool.Wait()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error occured on server shutting down: %s", err.Error())
	}
}
