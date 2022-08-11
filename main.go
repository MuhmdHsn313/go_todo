package main

import "github.com/kataras/iris/v12"

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

	return app
}
