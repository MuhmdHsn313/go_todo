package repositories

import (
	"sample_rest_api/models"
	"sample_rest_api/parameters"
)

type TodoRepository interface {
	FindAll(filter parameters.FilterTodo) ([]*models.Todo, error)
	FindTodo(id uint)
	CreateTodo(params parameters.NewTodo) (*models.Todo, error)
	UpdateTodo(id uint, params parameters.UpdateTodo) (*models.Todo, error)
	DeleteTodo(id uint) error
	CompleteTodo(id uint) error
	UncompleteTodo(id uint) error
}
