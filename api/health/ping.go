package health

import (
	"context"

	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)

var _ctx = context.Background()

// Ping returns a simple Pong
func Ping(app *iris.Application, deps *dependencies.Dependencies) {
	app.Get("/ping", func(ctx iris.Context) {
		pong, _ := deps.DataSource.Redis.Ping(_ctx).Result()

		ctx.JSON(iris.Map{"something": "Hello", "pong": pong})
	})
}
