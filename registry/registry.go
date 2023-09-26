package registry

import (
	"sample_rest_api/controllers"

	"github.com/kataras/iris/v12/middleware/jwt"
	"gorm.io/gorm"
)

type Registry struct {
	db     *gorm.DB
	signer *jwt.Signer
}

func NewRegistry(db *gorm.DB, signer *jwt.Signer) Registry {
	return Registry{db: db, signer: signer}
}

func (r Registry) NewAppController() controllers.AppContoller {
	return controllers.AppContoller{
		UserController: r.NewUserController(),
		TodoController: r.NewTodoController(),
	}
}
