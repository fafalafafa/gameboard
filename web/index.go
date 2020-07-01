package web

import (
	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)


func hostJS(app *iris.Application, dependencies *dependencies.Dependencies) {
	app.HandleDir("/js", "./public/js", iris.DirOptions{
		ShowList: true,
		Asset:      Asset,
		AssetInfo:  AssetInfo,
		AssetNames: AssetNames,
	})

}

func hostHTML(app *iris.Application, dependencies *dependencies.Dependencies) {
	tmpl := iris.HTML("./web/public", ".html")
	app.RegisterView(tmpl)
	app.Get("/*", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hi.html
		ctx.View("index.html")
	})
}
// New creates an instance of web
func New(app *iris.Application, dependencies *dependencies.Dependencies) {
	hostJS(app, dependencies)
	hostHTML(app, dependencies)
}
