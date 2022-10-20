package middleware

import (
	"github.com/sujit-baniya/framework/contracts/http"
)

func Test() http.HandlerFunc {
	return func(ctx http.Context) error {
		return ctx.Next()
	}
}
