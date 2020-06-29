package main

import (
	"context"

	"github.com/fafalafafa/gameboard/dependencies"

	"github.com/fafalafafa/gameboard/api"
	"github.com/fafalafafa/gameboard/config"
	"github.com/fafalafafa/gameboard/datasource"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

var ctx = context.Background()

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())

	app.Use(logger.New())

	DS := datasource.InitDS()

	CONFIG := config.New()

	DS.Redis.Ping(ctx).Result()

	Dependencies := dependencies.Initialize(DS, CONFIG)

	api.Initialize(app, Dependencies)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
