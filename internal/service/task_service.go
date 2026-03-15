package service

import (
	"task-tracker/internal/model"
	"task-tracker/internal/repository"
)

type TaskService struct{
	repo *repository.Repository
}

//functions
func NewTaskService(r *repository.Repository) *TaskService{
	return &TaskService{repo: r}
}

func (s *TaskService) CreateTask(t *model.Task) error{
	return s.repo.CreateTask(t)
}

func (s *TaskService) GetTasks() ([]*model.Task, error){
	return s.repo.GetTasks()
}

func (s *TaskService) GetTask(id int) (*model.Task, error){
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(t *model.Task) error{
	return s.repo.UpdateTask(t)
}

func (s *TaskService) DeleteTask(id int) error{
	return s.repo.DeleteTask(id)
}
