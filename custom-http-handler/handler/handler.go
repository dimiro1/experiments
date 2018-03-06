package handler

import (
	"github.com/dimiro1/experiments/custom-http-handler/binder"
	"github.com/dimiro1/experiments/custom-http-handler/params"
	"github.com/dimiro1/experiments/custom-http-handler/render"
	"github.com/dimiro1/experiments/custom-http-handler/validator"
)

type Default struct {
	params.Parameters
	binder.Binder
	validator.Validator
	render.Renderer
}
