package converter

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/auto_transcription/src/domains/converter"
)

type ConvertSpeechToTextUploadFileAction interface {
	Invoke() echo.HandlerFunc
}

type convertSpeechToTextUploadFileAction struct {
	useCase    converter.SpeechToTextUseCase
	translator converter.SpeechToTextTranslator
	responder  ConvertSpeechToTextUploadFileResponder
}

func NewConvertSpeechToTextUploadFileAction(
	useCase converter.SpeechToTextUseCase,
	translator converter.SpeechToTextTranslator,
	responder ConvertSpeechToTextUploadFileResponder,
) ConvertSpeechToTextUploadFileAction {
	return &convertSpeechToTextUploadFileAction{useCase, translator, responder}
}

func (a *convertSpeechToTextUploadFileAction) Invoke() echo.HandlerFunc {
	return func(context echo.Context) error {
		criteria, err := a.translator.EchoContextToCriteria(context)
		if err != nil {
			return a.responder.Invoke(context, nil, err)
		}
		result, err := a.useCase.UploadAndConvert(criteria)
		if err != nil {
			return a.responder.Invoke(context, nil, err)
		}
		return a.responder.Invoke(context, result, err)
	}
}
