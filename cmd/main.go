// @title Task Tracker API
// @version 1.0
// @description API для управления задачами
// @host localhost:8080
// @BasePath /
package main

import (
	"context"
	"log"
	_ "task-tracker/docs"
	"task-tracker/internal/handler"
	"task-tracker/internal/repository"
	"task-tracker/internal/service"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	if err := repository.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewTaskService(repo)
	handler := handler.NewHandler(service)

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Static("/", "web")
	e.File("/", "web/index.html")
	e.POST("/tasks", handler.CreateTask)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/tasks", handler.GetTasks)
	e.GET("/tasks/:id", handler.GetTask)
	e.PUT("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
