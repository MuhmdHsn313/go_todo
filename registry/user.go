package registry

import (
	"sample_rest_api/controllers"
	"sample_rest_api/repositories"
)

func (r Registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserRepository(), r.NewSessionRepository())
}

func (r Registry) NewUserRepository() repositories.UserRepository {
	return repositories.NewUserRepository(r.db)
}

func (r Registry) NewSessionRepository() repositories.SessionRepository {
	return repositories.NewSessionRepo(r.db, r.signer)
}
