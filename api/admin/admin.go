package admin

import (
	"context"

	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)

var _ctx = context.Background()

// ListCharacters returns a simple Pong
func ListCharacters(app *iris.Application, deps *dependencies.Dependencies) {
	app.Get("/admin/get-all-characters", func(ctx iris.Context) {
		characters, _ := deps.DataSource.Redis.LRange(_ctx, "characternames", 0, 1000).Result()

		ctx.JSON(iris.Map{"characters": characters})
	})
}
