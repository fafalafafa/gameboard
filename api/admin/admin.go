package admin

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)

var _ctx = context.Background()

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// ListCharacters returns a simple Pong
func ListCharacters(app *iris.Application, deps *dependencies.Dependencies) {
	app.Get("/admin/get-all-characters", func(ctx iris.Context) {
		characters, _ := deps.DataSource.Redis.LRange(_ctx, "characternames", 0, 1000).Result()

		ctx.JSON(iris.Map{"characters": characters})
	})

	app.Get("/admin/start-game", func(ctx iris.Context) {
		deps.DataSource.Redis.FlushAll(_ctx).Result()
		characters := []string{
			"mafia",
			"mafia",
			"doctor",
			"investigator-1",
			"investigator-2",
			"civilian",
			"civilian",
			"civilian",
			"civilian",
			"civilian",
		}

		length := len(characters)

		for i := 0; i < length; i++ {
			item := rand.Intn(len(characters))
			character := characters[item]
			characters = removeIndex(characters, item)
			fmt.Println(character)
			deps.DataSource.Redis.RPush(_ctx, "characters", character).Result()

		}

		ctx.JSON(iris.Map{"characters": "oink"})
	})
}
