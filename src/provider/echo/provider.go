package echo

import (
	"github.com/pt-suzuki/auto_transcription/src/provider/echo/controllers"
	"github.com/pt-suzuki/auto_transcription/src/provider/echo/middlewares"
)

type Provider struct {
	ControllerProvider *controllers.ControllerProvider
	MiddlewareProvider *middlewares.MiddlewareProvider
}

func NewProvider(
	controllerProvider *controllers.ControllerProvider,
	middlewareProvider *middlewares.MiddlewareProvider,
) *Provider {
	return &Provider{
		controllerProvider,
		middlewareProvider,
	}
}
