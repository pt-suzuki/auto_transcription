package echo

import (
	"github.com/pt-suzuki/auto_transcription/src/controllers/echo/converter"
	"github.com/pt-suzuki/auto_transcription/src/middleware"
)

type Provider struct {
	ConvertSpeechToTextAction       converter.ConvertSpeechToTextAction
	FirebaseTokenVerifiedMiddleware middleware.FirebaseTokenVerifiedMiddleware
}

func NewProvider(
	convertSpeechToTextAction converter.ConvertSpeechToTextAction,
	firebaseTokenVerifiedMiddleware middleware.FirebaseTokenVerifiedMiddleware,
) *Provider {
	return &Provider{
		convertSpeechToTextAction,
		firebaseTokenVerifiedMiddleware,
	}
}
