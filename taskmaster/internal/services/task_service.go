package services
// Interfsce for defining what a task can do

import "taskmaster/internal/models"


type TaskService interface {
	CreateTask(t *models.Task) error
	GetTask(id int64) (*models.Task, error)
	CompleteTask(id int64) error
}