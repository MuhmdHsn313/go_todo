package router

import (
	"sample_rest_api/controllers"

	"github.com/kataras/iris/v12/core/router"
)

func SetupTodoRouter(party router.Party, controller controllers.TodoController) {
	todo := party.Party("/todo")
	todo.Post("", controller.CreateTodo)
	todo.Patch("/{id:uint}", controller.UpdateTodo)
}
