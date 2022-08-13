package router

import (
	"sample_rest_api/controllers"

	"github.com/kataras/iris/v12/core/router"
)

func SetupRouters(apiBuilder *router.APIBuilder, contoller controllers.AppContoller) {
	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowedMethods:   []string{"*"},
	// })
	// apiBuilder.Use(crs)

	api := apiBuilder.Party("/api")
	SetupUserRoutes(api, contoller)
	SetupTodoRouter(api, contoller)
}
