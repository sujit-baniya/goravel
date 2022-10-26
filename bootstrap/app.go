package bootstrap

import (
	"github.com/sujit-baniya/framework/foundation"
	"goravel/config"
)

func Boot() {
	config.Boot()

	app := foundation.Application{Providers: config.Providers()}
	app.Boot()
}
