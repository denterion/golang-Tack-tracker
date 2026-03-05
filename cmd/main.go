package main

import (
	"context"
	"fmt"
	"log"
	"task-tracker/internal/model"
	"task-tracker/internal/repository"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")

	defer db.Close(context.Background())

	repo := repository.NewRepository(db)

	task := &model.Task{
		Title:       "Моя первая задача",
		Description: "Просто тест",
		Status:      "todo",
	}

	err = repo.CreateTask(task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создана первая задача", task)
}
