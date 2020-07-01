package characters

import (
	"context"
	"fmt"
	"math/rand"
	"time"

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

		session := generateRandomString(16)
		fmt.Println("sessions:" + session)
		fmt.Println(character)
		_, err2 := deps.DataSource.Redis.Set(_ctx, "sessions:"+session, character, 99999999).Result()

		if err2 != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err2.Error())
			return
		}
		ctx.JSON(iris.Map{"character": character, "sessId": session})
	})
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generateRandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
