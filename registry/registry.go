package registry

import (
	"gorm.io/gorm"
	"sample_rest_api/controllers"
)

type Registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) Registry {
	return Registry{db: db}
}

func (r Registry) NewAppController() controllers.AppContoller {
	return controllers.AppContoller{
		UserController: r.NewUserController(),
	}
}
