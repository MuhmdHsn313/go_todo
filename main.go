package main

import (
	"github.com/kataras/iris/v12"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sample_rest_api/datastore"
	"sample_rest_api/registry"
	"sample_rest_api/router"
)

func main() {
	app := newApp()
	err := app.Listen(":8080")
	if err != nil {
		print(err.Error())
		return
	}
}

func newApp() *iris.Application {
	app := iris.New()

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = datastore.Mirgate(db)
	if err != nil {
		panic(err)
	}

	reg := registry.NewRegistry(db)
	router.SetupRouters(app.APIBuilder, reg.NewAppController())

	return app
}
