package routes

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/contracts/route"
	"github.com/sujit-baniya/middleware"
	"goravel/app/http/controllers"
)

func Web(route route.Route) {
	userController := controllers.NewUserController()
	route.Get("/", func(ctx http.Context) error {
		ctx.AbortWithStatus(500)
		return ctx.Render("index", map[string]any{
			"title": "This is test page",
		}, "layouts/master")
	})

	route.Get("/users/{id}", middleware.BasicAuth(middleware.ConfigBasicAuth{
		Users: map[string]string{
			"john": "doe",
		},
	}), userController.Show)
}
