package echo

import (
	"github.com/pt-suzuki/auto_transcription/src/controllers/echo/converter"
)

type Provider struct {
	ConvertSpeechToTextAction converter.ConvertSpeechToTextAction
	MiddlewareProvider        *MiddlewareProvider
}

func NewProvider(
	convertSpeechToTextAction converter.ConvertSpeechToTextAction,
	middlewareProvider *MiddlewareProvider,
) *Provider {
	return &Provider{
		convertSpeechToTextAction,
		middlewareProvider,
	}
}
