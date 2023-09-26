package registry

import (
	"sample_rest_api/controllers"
	"sample_rest_api/repositories"
)

func (r Registry) NewTodoController() controllers.TodoController {
	return controllers.NewTodoController(r.NewTodoRepository())
}

func (r Registry) NewTodoRepository() repositories.TodoRepository {
	return repositories.NewTodoRepo(r.db)
}
