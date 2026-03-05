package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"task-tracker/internal/model"
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

	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := repo.GetTasks()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/tasks", func(c echo.Context) error {
		t := new(model.Task)
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		if err := repo.CreateTask(t); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, t)
	})

	e.PUT("/tasks/id", func(c echo.Context) error {
		id := c.Param("id")
		t := new(model.Task)
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		var taskID int
		fmt.Scanf(id, "%d", &taskID)
		t.ID = taskID

		if err := repo.UpdateTask(t); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, t)
	})

	e.DELETE("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		var taskID int
		fmt.Scanf(id, "%d", &taskID)
		if err := repo.DeleteTask(taskID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
	})

	log.Println("Sever started on :8080")
	e.Logger.Fatal(e.Start(":8080"))

}
