package dependencies

import (
	"github.com/fafalafafa/gameboard/config"
	"github.com/fafalafafa/gameboard/datasource"
)

// Dependencies type will be exported
type Dependencies struct {
	DataSource *datasource.DataSource
	Config     *config.Config
}

// Initialize Creates a dependency
func Initialize(DS *datasource.DataSource, config *config.Config) *Dependencies {
	return &Dependencies{
		DataSource: DS,
		Config:     config,
	}
}
