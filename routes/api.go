package routes

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/contracts/route"
)

func Api(route route.Route) {
	route.Get("/", func(ctx http.Context) error {
		return ctx.Response().Json(200, http.Json{
			"status": "API Request",
		})
	})
}
