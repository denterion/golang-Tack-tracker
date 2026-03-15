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

// @title Task Tracker API
// @version 1.0
// @description API для управления задачами
// @host localhost:8080
// @BasePath /

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
	taskService := service.NewTaskService(repo)
	h := handler.NewHandler(taskService)

	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogError:   true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				c.Logger().Error(v.Error)
			}

			c.Logger().Infof(
				"%s %s %d %v",
				v.Method,
				v.URI,
				v.Status,
				v.Latency,
			)

			return nil
		},
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.Static("/static", "web/static")
	e.Static("/images", "web")
	
	e.File("/", "web/index.html")
	e.File("/about", "web/about.html")

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/tasks", h.CreateTask)
	e.GET("/tasks", h.GetTasks)
	e.GET("/tasks/:id", h.GetTask)
	e.PUT("/tasks/:id", h.UpdateTask)
	e.DELETE("/tasks/:id", h.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
