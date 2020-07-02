package admin

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/fafalafafa/gameboard/plugins"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
)

var _ctx = context.Background()

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

type Members struct {
	Name      string
	Character string
}

// ListCharacters returns a simple Pong
func ListCharacters(app *iris.Application, deps *dependencies.Dependencies) {
	app.Get("/admin/get-all-characters", func(ctx iris.Context) {
		characters, _ := deps.DataSource.Redis.LRange(_ctx, "sessions", 0, 1000).Result()
		gameID := getGame(deps.DataSource.Redis)
		var users []Members

		for i := 0; i < len(characters); i++ {
			var _user Members
			sessionID := characters[i]
			Name, _ := deps.DataSource.Redis.Get(_ctx, "character:"+gameID+":"+sessionID+":name").Result()
			Character, _ := deps.DataSource.Redis.Get(_ctx, "character:"+gameID+":"+sessionID+":character").Result()
			_user.Name = Name
			_user.Character = Character
			fmt.Println(_user.Name)
			fmt.Println(_user.Character)
			users = append(users, _user)
		}

		ctx.JSON(iris.Map{"characters": users})
	})

	app.Get("/admin/start-game", func(ctx iris.Context) {
		gameID := plugins.GenerateRandomString(16)
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
			deps.DataSource.Redis.RPush(_ctx, "characters", character).Result()

		}

		deps.DataSource.Redis.Set(_ctx, "game:id", gameID, 0).Result()

		ctx.JSON(iris.Map{"gameID": gameID})
	})
}

func getGame(Redis *redis.Client) string {
	gameID, err := Redis.Get(_ctx, "game:id").Result()

	if err == redis.Nil {
		return ""
	}

	return gameID
}
