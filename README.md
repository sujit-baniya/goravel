<p align="center"><img src="https://goravel.s3.us-east-2.amazonaws.com/goravel-word.png" width="300"></p>

Forked https://github.com/goravel/framework
With support to switch between [Gin](https://gin-gonic.com/), [Chi](https://github.com/sujit-baniya/chi) and [GoFiber](https://gofiber.io/)

## Switch between Gin or GoFiber
In `config/app.go`, in ServiceProvider list change `route.ServiceProvider` accordingly as

```go
    "providers": []contracts.ServiceProvider{
        ...
		// &route.ServiceProvider{Engine: route.NewFiber()},
		// &route.ServiceProvider{Engine: route.NewGin()},
		// &route.ServiceProvider{Engine: route.NewChi()},
		&route.ServiceProvider{}, // Default Chi (standard net/http) is used: https://github.com/sujit-baniya/chi
        ...
    },
```

## About Goravel

Goravel is a web application framework with complete functions and good scalability. As a starting scaffolding to help
Golang developers quickly build their own applications.

> IMPORTANT：Goravel v1 has been greatly upgraded and refactored, not compatible with v0, the v0 version is no longer
> updated and maintained, but its documentation can be found on [v0](https://github.com/goravel/docs/tree/master/v0)
> .

## Main Function

- [x] Config
- [x] Http
- [x] Orm
- [x] Migrate
- [x] Logger
- [x] Cache
- [x] Grpc
- [x] Artisan Console
- [x] Task Scheduling
- [x] Queue
- [x] Event
- [x] Mail
- [x] Mock

## Roadmap

- [ ] Optimize experience of microservice
- [ ] Orm relationships
- [ ] Request validator
- [ ] Jwt

## Documentation

Online documentation [https://www.goravel.dev](https://www.goravel.dev)

> To optimize the documentation, please submit a PR to the documentation
> repository [https://github.com/goravel/docs](https://github.com/goravel/docs)

## Group

Welcome more exchanges in Discord.

[https://discord.gg/cFc5csczzS](https://discord.gg/cFc5csczzS)

## Tribute Laravel

Goravel and Laravel remain highly consistent, let PHPer play Golang happily without learning a new framework!

## License

The Goravel framework is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
