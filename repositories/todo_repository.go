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
	FindTodo(id uint) (*models.Todo, error)
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
	tx := td.db.Limit(filter.GetLimit()).Offset(filter.GetOffest()).Order(filter.OrderQueryBy())
	if filter.IsDone != nil {
		tx.Where("is_done = ?", *filter.IsDone)
	}
	if filter.Title != nil {
		tx.Where("title LIKE ?", "%"+*filter.Title+"%")
	}
	err := tx.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (td *todoRepository) FindTodo(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (td *todoRepository) CreateTodo(params parameters.NewTodo) (*models.Todo, error) {
	todo := models.Todo{
		Title: params.Title,
		Body:  params.Body,
	}
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
	if err := td.db.Updates(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (td *todoRepository) DeleteTodo(id uint) error {
	return td.db.Delete(&models.Todo{}, id).Error
}

func (td *todoRepository) CompleteTodo(id uint) error {
	var todo models.Todo
	if err := td.db.First(&todo, id).Error; err != nil {
		return err
	}
	todo.IsDone = true
	if err := td.db.Updates(&todo).Error; err != nil {
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
	if err := td.db.Updates(&todo).Error; err != nil {
		return err
	}
	return nil
}
