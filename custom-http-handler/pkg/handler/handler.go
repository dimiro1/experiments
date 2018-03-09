package handler

import (
	"github.com/dimiro1/experiments/custom-http-handler/pkg/binder"
	"github.com/dimiro1/experiments/custom-http-handler/pkg/params"
	"github.com/dimiro1/experiments/custom-http-handler/pkg/render"
	"github.com/dimiro1/experiments/custom-http-handler/pkg/validator"
)

type Default struct {
	params.Parameters
	binder.Binder
	validator.Validator
	render.Renderer
}
