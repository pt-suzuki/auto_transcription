package echo

import (
	"github.com/pt-suzuki/auto_transcription/src/middleware"
)

type MiddlewareProvider struct {
	FirebaseTokenVerifiedMiddleware middleware.FirebaseTokenVerifiedMiddleware
	ConverterMiddlewareProvider     *ConverterMiddlewareProvider
}

func NewMiddlewareProvider(
	firebaseTokenVerifiedMiddleware middleware.FirebaseTokenVerifiedMiddleware,
	converterMiddlewareProvider *ConverterMiddlewareProvider,
) *MiddlewareProvider {
	return &MiddlewareProvider{
		firebaseTokenVerifiedMiddleware,
		converterMiddlewareProvider,
	}
}
