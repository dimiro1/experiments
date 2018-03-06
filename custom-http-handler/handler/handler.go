package handler

import (
	"github.com/dimiro1/experiments/custom-http-handler/binder"
	"github.com/dimiro1/experiments/custom-http-handler/params"
	"github.com/dimiro1/experiments/custom-http-handler/render"
	"github.com/dimiro1/experiments/custom-http-handler/validator"
)

type Default struct {
	Params    params.Parameters
	Binder    binder.Binder
	Validator validator.Validator
	Renderer  render.Renderer
}
