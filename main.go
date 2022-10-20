package main

import (
	"github.com/sujit-baniya/framework/facades"

	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start http server by facades.Route.
	go func() {
		if err := facades.Route.Run(facades.Config.GetString("app.host")); err != nil {
			facades.Log.Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
