package http

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/middleware"
)

type Kernel struct{}

func (kernel *Kernel) Middleware() []http.HandlerFunc {
	return []http.HandlerFunc{
		middleware.Secure(),
		middleware.Recover(),
		/*middleware.Log(middleware.ConfigLog{
			LogWriter: &log.FileWriter{
				Filename:     "logs/main.log",
				FileMode:     0600,
				MaxSize:      100 * 1024 * 1024,
				MaxBackups:   7,
				TimeFormat:   "2006-01-02",
				EnsureFolder: true,
				LocalTime:    false,
			},
		}),*/
		// appMiddleware.Test(),
	}
}

func (kernel *Kernel) Web() []http.HandlerFunc {
	return []http.HandlerFunc{
		middleware.Cors(),
	}
}

func (kernel *Kernel) Api() []http.HandlerFunc {
	return []http.HandlerFunc{}
}
