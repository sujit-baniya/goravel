package http

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/http/middleware"
)

type Kernel struct {
}

func (kernel *Kernel) Middleware() []http.HandlerFunc {
	return []http.HandlerFunc{
		middleware.Cors(),
		// appMiddleware.Test(),
	}
}
