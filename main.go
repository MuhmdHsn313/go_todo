package main

import "github.com/kataras/iris/v12"

func main() {
	app := newApp()
	app.Listen(":8080")
}

func newApp() *iris.Application {
	app := iris.New()

	return app
}
