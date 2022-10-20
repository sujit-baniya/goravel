package bootstrap

import (
	"github.com/sujit-baniya/framework/foundation"
	"goravel/config"
)

func Boot() {
	app := foundation.Application{}

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()
}
