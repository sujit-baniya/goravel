package providers

import (
	"github.com/sujit-baniya/framework/facades"
	"goravel/app/http"
	"goravel/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Boot() {

}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares.
	kernel := http.Kernel{}
	facades.Route.GlobalMiddleware(kernel.Middleware()...)

	//Add web routes
	webRoutes := facades.Route.Middleware(kernel.Web()...)
	routes.Web(webRoutes)

	//Add api routes
	apiRoutes := facades.Route.
		Prefix(facades.Config.GetString("api.prefix") + "/" + facades.Config.GetString("api.version")).
		Middleware(kernel.Api()...)

	routes.Api(apiRoutes)
}
