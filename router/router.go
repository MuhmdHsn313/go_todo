package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/core/router"
	"sample_rest_api/controllers"
)

func SetupRouters(apiBuilder *router.APIBuilder, contoller controllers.AppContoller) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"*"},
	})
	apiBuilder.Use(crs)

	api := apiBuilder.Party("/api")
	SetupUserRoutes(api, contoller)
}
