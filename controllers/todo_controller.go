package controllers

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type todoController struct {
	db *gorm.DB
}

type TodoController interface {
	CreateTodo(ctx iris.Context)
}

func NewTodoController(db *gorm.DB) TodoController {
	return &todoController{db}
}

func (tc *todoController) CreateTodo(ctx iris.Context) {

}
