package server

import (
	"github.com/pt-suzuki/auto_transcription/src/handler"
)

func ProviderResponseHandler() handler.ResponseHandler {
	return handler.NewResponseHandler()
}
