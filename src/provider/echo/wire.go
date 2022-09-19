//go:build wireinject
// +build wireinject

package echo

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"github.com/google/wire"
	"github.com/pt-suzuki/auto_transcription/src/controllers/echo/converter"
	converter2 "github.com/pt-suzuki/auto_transcription/src/domains/converter"
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
	"github.com/pt-suzuki/auto_transcription/src/handler"
	"github.com/pt-suzuki/auto_transcription/src/middleware"
	converter3 "github.com/pt-suzuki/auto_transcription/src/middleware/converter"
)

func Wire(fireStoreClient *firestore.Client, fireStorageClient *storage.Client, firebaseAuthClient *auth.Client) *Provider {
	wire.Build(
		handler.NewResponseHandler,
		uploader.NewTranslator,
		uploader.NewRepository,
		uploader.NewUseCase,
		middleware.NewFirebaseTokenVerifiedMiddleware,
		converter2.NewSpeechToTextTranslator,
		converter2.NewSpeechToTextRepository,
		converter2.NewSpeechToTextUseCase,
		converter2.NewConvertResultTranslator,
		converter2.NewConvertResultRepository,
		converter2.NewConvertResultUseCase,
		converter.NewConvertSpeechToTextResponder,
		converter.NewConvertSpeechToTextAction,
		converter3.NewConvertSpeechToTextValidatorMiddleware,
		NewConverterMiddlewareProvider,
		NewMiddlewareProvider,
		NewProvider,
	)
	return &Provider{}
}
