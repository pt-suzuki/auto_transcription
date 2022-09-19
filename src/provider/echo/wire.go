// +build wireinject

package echo

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/storage"
	"github.com/google/wire"
	"github.com/pt-suzuki/auto_transcription/src/controllers/echo/converter"
	converter2 "github.com/pt-suzuki/auto_transcription/src/domains/converter"
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
	"github.com/pt-suzuki/auto_transcription/src/handler"
)

func Wire(fireStoreClient *firestore.Client, fireStorageClient *storage.Client) *Provider {
	wire.Build(
		handler.NewResponseHandler,
		uploader.NewTranslator,
		uploader.NewRepository,
		uploader.NewUseCase,
		converter2.NewSpeechToTextTranslator,
		converter2.NewSpeechToTextRepository,
		converter2.NewSpeechToTextUseCase,
		converter2.NewConvertResultTranslator,
		converter2.NewConvertResultRepository,
		converter2.NewConvertResultUseCase,
		converter.NewConvertSpeechToTextResponder,
		converter.NewConvertSpeechToTextAction,
		NewProvider,
	)
	return &Provider{}
}
