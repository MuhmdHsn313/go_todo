package controllers

import (
	"sample_rest_api/parameters"
	"sample_rest_api/repositories"

	"github.com/kataras/iris/v12"
)

type todoController struct {
	TodoRepository repositories.TodoRepository
}

type TodoController interface {
	CreateTodo(ctx iris.Context)
	UpdateTodo(ctx iris.Context)
}

func NewTodoController(ur repositories.TodoRepository) TodoController {
	return &todoController{TodoRepository: ur}
}

func (tc *todoController) CreateTodo(ctx iris.Context) {
	var params parameters.NewTodo
	err := ctx.ReadBody(&params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	todo, err := tc.TodoRepository.CreateTodo(params)
	if err != nil {
		return
	}
	ctx.StopWithJSON(iris.StatusCreated, todo)

}

func (tc *todoController) UpdateTodo(ctx iris.Context) {
	var params parameters.UpdateTodo
	id, _ := ctx.Params().GetUint("id")
	err := ctx.ReadBody(&params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	todo, err := tc.TodoRepository.UpdateTodo(id, params)
	if err != nil {
		return
	}
	ctx.StopWithJSON(iris.StatusOK, todo)
}
