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

	tasks, err := repo.GetTasks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Список всех задач: ")
	for _, t := range tasks {
		fmt.Println(t)
	}

	task.Title = "Обнова"
	task.Description = "Уху"
	task.Status = "done"

	err = repo.UpdateTask(task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Задача Обновлена!")

	tasks, err = repo.GetTasks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Список после обновления")

	for _, t := range tasks {
		fmt.Println(t)
	}

	err = repo.DeleteTask(task.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Задача удалена!")

	tasks, err = repo.GetTasks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Список после удаления")
	for _, t := range tasks {
		fmt.Println(t)
	}
}
