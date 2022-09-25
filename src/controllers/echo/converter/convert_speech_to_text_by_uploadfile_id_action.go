package converter

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/auto_transcription/src/domains/converter"
)

type ConvertSpeechToTextByUploadFileIDAction interface {
	Invoke() echo.HandlerFunc
}

type convertSpeechToTextByUploadFileIDAction struct {
	useCase    converter.SpeechToTextUseCase
	translator converter.SpeechToTextTranslator
	responder  ConvertSpeechToTextByUploadFileIDResponder
}

func NewConvertSpeechToTextByUploadFileIDAction(
	useCase converter.SpeechToTextUseCase,
	translator converter.SpeechToTextTranslator,
	responder ConvertSpeechToTextByUploadFileIDResponder,
) ConvertSpeechToTextByUploadFileIDAction {
	return &convertSpeechToTextByUploadFileIDAction{useCase, translator, responder}
}

func (a *convertSpeechToTextByUploadFileIDAction) Invoke() echo.HandlerFunc {
	return func(context echo.Context) error {
		id := a.translator.EchoContextToId(context)
		result, err := a.useCase.ConvertByUploadFileId(id)
		if err != nil {
			return a.responder.Invoke(context, nil, err)
		}
		return a.responder.Invoke(context, result, err)
	}
}
