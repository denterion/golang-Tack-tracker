package main

import (
	"context"
	"log"
	"task-tracker/internal/handler"
	"task-tracker/internal/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	repo := repository.NewRepository(db)
	h := handler.NewHandler(repo)

	e := echo.New()
	e.Static("/", "web")

	e.POST("/tasks", h.CreateTask)
	e.GET("/tasks", h.GetTasks)
	e.GET("/tasks/:id", h.GetTask)
	e.PUT("/tasks/:id", h.UpdateTask)
	e.DELETE("/tasks/:id", h.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
