package controllers

import (
	"github.com/sujit-baniya/framework/contracts/http"
	"github.com/sujit-baniya/framework/facades"
	"goravel/app/jobs"
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
	facades.Queue.Job(&jobs.SendEmails{}, nil).OnQueue("default").Dispatch()
	return ctx.Json(http.Json{
		"Hello": ctx.Params("id"),
	})
}
