package bootstrap

import (
	"github.com/sujit-baniya/framework/foundation"
	"goravel/config"
)

func Boot() {
	app := foundation.Application{}

	//Bootstrap the config.
	config.Boot()

	//Bootstrap the application
	app.Boot()
}
