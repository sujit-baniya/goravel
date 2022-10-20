package routes

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/facades"
	"goravel/app/http/middleware"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) error {
		return ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

	userController := controllers.NewUserController()
	facades.Route.Get("/users/{id}", middleware.Test(), userController.Show)
}
