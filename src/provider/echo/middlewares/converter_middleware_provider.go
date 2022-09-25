package middlewares

import (
	"github.com/pt-suzuki/auto_transcription/src/middleware/converter"
)

type ConverterMiddlewareProvider struct {
	ConvertSpeechToTextValidatorMiddleware converter.ConvertSpeechToTextValidatorMiddleware
}

func NewConverterMiddlewareProvider(
	ConvertSpeechToTextValidatorMiddleware converter.ConvertSpeechToTextValidatorMiddleware,
) *ConverterMiddlewareProvider {
	return &ConverterMiddlewareProvider{
		ConvertSpeechToTextValidatorMiddleware,
	}
}
