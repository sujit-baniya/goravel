package routes

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/contracts/route"
	"github.com/sujit-baniya/framework/facades"
	"github.com/sujit-baniya/middleware"
	"goravel/app/http/controllers"
)

func Web(route route.Route) {
	userController := controllers.NewUserController()
	route.Get("/", func(ctx http.Context) error {
		if !facades.Cache.Has("test1") {
			facades.Cache.Forever("test1", "Awesome!!!")
		}
		return ctx.Render("index", map[string]any{
			"title": facades.Cache.GetString("test1", "Not found"),
		}, "layouts/master")
	})

	route.Get("/users/{id}", middleware.BasicAuth(middleware.ConfigBasicAuth{
		Users: map[string]string{
			"john": "doe",
		},
	}), userController.Show)
}
