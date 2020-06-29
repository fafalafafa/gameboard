package web

import (
	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)

// New creates an instance of web
func New(app *iris.Application, dependencies *dependencies.Dependencies) {
	tmpl := iris.HTML("./html", ".html")

	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hi.html
		ctx.View("index.html")
	})
}
