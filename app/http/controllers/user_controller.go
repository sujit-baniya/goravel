package controllers

import (
	"github.com/sujit-baniya/framework/contracts/http"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) error {
	return ctx.Response().Success().Json(http.Json{
		"Hello": ctx.Params("id"),
	})
}