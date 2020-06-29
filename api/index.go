package api

import (
	"github.com/fafalafafa/gameboard/api/admin"
	"github.com/fafalafafa/gameboard/api/characters"
	"github.com/fafalafafa/gameboard/api/health"
	"github.com/fafalafafa/gameboard/dependencies"
	"github.com/kataras/iris/v12"
)

// Initialize defines all the API urls that will be provided by the application
func Initialize(app *iris.Application, dependencies *dependencies.Dependencies) {
	health.Ping(app, dependencies)
	admin.ListCharacters(app, dependencies)
	characters.InitializeCharacter(app, dependencies)
}
