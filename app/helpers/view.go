package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/fiberRoute"
	contractRoute "github.com/sujit-baniya/framework/contracts/route"
	"github.com/sujit-baniya/framework/facades"
	"github.com/sujit-baniya/framework/http"
	"github.com/sujit-baniya/framework/route"
	"github.com/sujit-baniya/framework/view"
	"github.com/sujit-baniya/ginRoute"
)

func GetRouter() contractRoute.Engine {
	router := facades.Config.GetString("app.router")
	template := facades.Config.GetString("view.template")
	extension := facades.Config.GetString("view.extension")
	//Default Chi (standard net/http) is used: https://github.com/sujit-baniya/chi
	viewEngine := view.New(template, extension)
	switch router {
	case "fiber":
		return fiberRoute.New(fiber.Config{
			Views: viewEngine,
		})
	case "gin":
		return ginRoute.New(ginRoute.Config{
			View: viewEngine,
		})
	default:
		return route.NewChi(http.ChiConfig{
			View: viewEngine,
		})
	}
}
