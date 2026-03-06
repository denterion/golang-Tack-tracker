package handler

import (
	"net/http"
	"strconv"
	"task-tracker/internal/model"
	"task-tracker/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct{
	repo *repository.Repository
}

func NewHandler(r * repository.Repository) *Handler{
	return &Handler{repo: r}
}

func (h *Handler) CreateTask(c echo.Context) error{
	t := new(model.Task)

	if err := c.Bind(t); err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.repo.CreateTask(t)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, t)
}

func (h *Handler) GetTasks(c echo.Context) error{
	
	tasks, err := h.repo.GetTasks()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTask(c echo.Context) error{

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	task, err := h.repo.GetTaskByID(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, task)
}

func (h *Handler) UpdateTask(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	t := new(model.Task)

	if err := c.Bind(t); err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	t.ID = id

	err = h.repo.UpdateTask(t)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, t)
}

func (h *Handler) DeleteTask(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	err = h.repo.DeleteTask(id)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "task deleted",
	})
}