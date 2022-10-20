package config

import (
	"github.com/sujit-baniya/framework/cache"
	"github.com/sujit-baniya/framework/console"
	"github.com/sujit-baniya/framework/contracts"
	"github.com/sujit-baniya/framework/database"
	"github.com/sujit-baniya/framework/event"
	"github.com/sujit-baniya/framework/facades"
	"github.com/sujit-baniya/framework/grpc"
	"github.com/sujit-baniya/framework/http"
	"github.com/sujit-baniya/framework/log"
	"github.com/sujit-baniya/framework/mail"
	"github.com/sujit-baniya/framework/queue"
	"github.com/sujit-baniya/framework/route"
	"github.com/sujit-baniya/framework/schedule"

	"goravel/app/providers"
)

// Boot Start all init methods of the current folder to bootstrap all config.
func Boot() {}

func init() {
	config := facades.Config
	config.Add("app", map[string]interface{}{
		//Application Name
		//This value is the name of your application. This value is used when the
		//framework needs to place the application's name in a notification or
		//any other location as required by the application or its packages.
		"name": config.Env("APP_NAME", "Goravel"),

		//Application Environment
		//This value determines the "environment" your application is currently
		//running in. This may determine how you prefer to configure various
		//services the application utilizes. Set this in your ".env" file.
		"env": config.Env("APP_ENV", "production"),

		//Application Debug Mode
		"debug": config.Env("APP_DEBUG", false),

		//Application Timezone
		//Here you may specify the default timezone for your application, which
		//will be used by the PHP date and date-time functions. We have gone
		//ahead and set this to a sensible default for you out of the box.
		"timezone": "UTC",

		//Encryption Key
		//32 character string, otherwise these encrypted strings
		//will not be safe. Please do this before deploying an application!
		"key": config.Env("APP_KEY", ""),

		//Application URL
		"url": config.Env("APP_URL", "http://localhost"),

		//Application host, http server listening address.
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),

		"grpc_host": config.Env("GRPC_HOST", ""),

		//Autoload service providers
		//The service providers listed here will be automatically loaded on the
		//request to your application. Feel free to add your own services to
		//this array to grant expanded functionality to your applications.
		"providers": []contracts.ServiceProvider{
			&log.ServiceProvider{},
			&console.ServiceProvider{},
			&database.ServiceProvider{},
			&cache.ServiceProvider{},
			&http.ServiceProvider{},
			// &route.ServiceProvider{Engine: route.NewFiber()},
			// &route.ServiceProvider{Engine: route.NewGin()},
			&route.ServiceProvider{}, // Default Gin is used
			&schedule.ServiceProvider{},
			&event.ServiceProvider{},
			&queue.ServiceProvider{},
			&grpc.ServiceProvider{},
			&mail.ServiceProvider{},
			&providers.AppServiceProvider{},
			&providers.RouteServiceProvider{},
			&providers.GrpcServiceProvider{},
			&providers.ConsoleServiceProvider{},
			&providers.QueueServiceProvider{},
			&providers.EventServiceProvider{},
		},
	})
}
