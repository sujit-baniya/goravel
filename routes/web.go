package routes

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/contracts/route"
	"github.com/sujit-baniya/framework/http/middleware"
	"github.com/sujit-baniya/framework/http/middleware/limiter"
	"goravel/app/http/controllers"
)

func Web(route route.Route) {
	userController := controllers.NewUserController()
	route.Get("/", middleware.RequestID(), limiter.New(), func(ctx http.Context) error {
		return ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

	route.Get("/users/{id}", middleware.BasicAuth(middleware.ConfigBasicAuth{
		Users: map[string]string{
			"john": "doe",
		},
	}), userController.Show)
}
