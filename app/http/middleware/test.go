package middleware

import (
	"fmt"
	"github.com/sujit-baniya/framework/contracts/http"
)

func Test() http.HandlerFunc {
	return func(ctx http.Context) error {
		fmt.Println("Test middleware")
		return ctx.Request().Next()
	}
}
