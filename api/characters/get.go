package characters

import (
	"context"

	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)

var _ctx = context.Background()

// User struct defines the user form
type User struct {
	Name string `form:"name"`
}

// InitializeCharacter defines the APIS for character
func InitializeCharacter(app *iris.Application, deps *dependencies.Dependencies) {
	app.Get("/get-character", func(ctx iris.Context) {
		character, _ := deps.DataSource.Redis.LPop(_ctx, "characters").Result()

		ctx.JSON(iris.Map{"something": "Hello", "pong": character})
	})

	app.Post("/get-character", func(ctx iris.Context) {
		var player User
		err := ctx.ReadForm(&player)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		// Set it
		character, err := deps.DataSource.Redis.LPop(_ctx, "characters").Result()

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		// if it did not have an error then set it to this.

		deps.DataSource.Redis.RPush(_ctx, "characternames", player.Name+"-"+character)

		ctx.Writef(character)
		// ctx.Writef("Received: %#+v\n", player)
		// ctx.WriteString(player.Name)
	})
}
