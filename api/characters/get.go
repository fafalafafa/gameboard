package characters

import (
	"context"
	"fmt"

	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/fafalafafa/gameboard/plugins"
	"github.com/go-redis/redis/v8"
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
		// character, _ := deps.DataSource.Redis.LPop(_ctx, "characters").Result()
		// Get the game ID first
		gameID := getGame(deps.DataSource.Redis)

		if gameID == "" {
			ctx.JSON(iris.Map{"message": "No Games Yet", "result": "failed"})
			return
		}

		sessionID := ctx.URLParamDefault("sessionId", "nothing")

		role, _ := deps.DataSource.Redis.Get(_ctx, "sessions:"+gameID+":"+sessionID).Result()
		// If there is no character
		if role == "" {
			fmt.Printf("Character for %s %s does not exist\n", gameID, sessionID)
			ctx.JSON(iris.Map{"message": nil, "result": "success"})
			return
		}

		ctx.JSON(iris.Map{"message": role, "result": "success"})
	})

	app.Post("/get-character", func(ctx iris.Context) {
		var player User
		err := ctx.ReadForm(&player)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		gameID := getGame(deps.DataSource.Redis)

		if gameID == "" {
			ctx.JSON(iris.Map{"message": "No Games Yet", "result": "failed"})
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

		session := plugins.GenerateRandomString(16)

		deps.DataSource.Redis.Set(_ctx, "character:"+gameID+":"+session+":name", player.Name, 0)
		deps.DataSource.Redis.Set(_ctx, "character:"+gameID+":"+session+":character", character, 0)
		deps.DataSource.Redis.RPush(_ctx, "sessions", session)

		_, err2 := deps.DataSource.Redis.Set(_ctx, "sessions:"+gameID+":"+session, character, 0).Result()

		if err2 != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err2.Error())
			return
		}

		ctx.JSON(iris.Map{"character": character, "sessId": session})
	})
}

func getGame(Redis *redis.Client) string {
	gameID, err := Redis.Get(_ctx, "game:id").Result()

	if err == redis.Nil {
		return ""
	}

	return gameID
}
