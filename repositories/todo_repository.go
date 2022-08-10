package repositories

import (
	"sample_rest_api/models"
	"sample_rest_api/parameters"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

type TodoRepository interface {
	FindAll(filter parameters.FilterTodo) ([]*models.Todo, error)
	FindTodo(id uint)
	CreateTodo(params parameters.NewTodo) (*models.Todo, error)
	UpdateTodo(id uint, params parameters.UpdateTodo) (*models.Todo, error)
	DeleteTodo(id uint) error
	CompleteTodo(id uint) error
	UncompleteTodo(id uint) error
}

func NewTodoRepo(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (td *todoRepository) FindAll(filter parameters.FilterTodo) ([]*models.Todo, error) {

	var todos []*models.Todo
	if err := td.db.Limit(filter.Limit).Offset(filter.Offset).Order(filter.OrderField+" "+filter.OrderBy).Where(
		"is_done = ?", filter.IsDone).Find(&todos).Error; err != nil {
		return nil, err

	}
	return todos, nil
}

func (td *todoRepository) FindTodo(id uint) {
	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return
	}
}

func (td *todoRepository) CreateTodo(params parameters.NewTodo) (*models.Todo, error) {
	var todo models.Todo
	todo.Title = params.Title
	todo.Body = params.Body
	if err := td.db.Create(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (td *todoRepository) UpdateTodo(id uint, params parameters.UpdateTodo) (*models.Todo, error) {
	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	if params.Title != nil {
		todo.Title = *params.Title
	}
	if params.Body != nil {
		todo.Body = *params.Body
	}
	if err := td.db.Save(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (td *todoRepository) DeleteTodo(id uint) error {
	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return err
	}
	if err := td.db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (td *todoRepository) CompleteTodo(id uint) error {

	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return err
	}
	todo.IsDone = true
	if err := td.db.Save(&todo).Error; err != nil {
		return err
	}
	return nil

}

func (td *todoRepository) UncompleteTodo(id uint) error {
	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return err
	}
	todo.IsDone = false
	if err := td.db.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}
