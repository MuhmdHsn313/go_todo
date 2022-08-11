package router

import (
	"github.com/kataras/iris/v12/core/router"
	"sample_rest_api/controllers"
)

func SetupUserRoutes(party router.Party, controller controllers.UserController) {
	user := party.Party("/user")

	user.Post("", controller.CreateUser)
	user.Post("/login", controller.LoginUser)
}
