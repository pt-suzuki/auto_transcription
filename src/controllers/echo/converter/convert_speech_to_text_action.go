package converter

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/auto_transcription/src/domains/converter"
)

type ConvertSpeechToTextAction interface {
	Invoke() echo.HandlerFunc
}

type convertSpeechToTextAction struct {
	useCase    converter.SpeechToTextUseCase
	translator converter.SpeechToTextTranslator
	responder  ConvertSpeechToTextResponder
}

func NewConvertSpeechToTextAction(
	useCase converter.SpeechToTextUseCase,
	translator converter.SpeechToTextTranslator,
	responder ConvertSpeechToTextResponder,
) ConvertSpeechToTextAction {
	return &convertSpeechToTextAction{useCase, translator, responder}
}

func (a *convertSpeechToTextAction) Invoke() echo.HandlerFunc {
	return func(context echo.Context) error {
		criteria, err := a.translator.EchoContextToCriteria(context)
		if err != nil {
			return a.responder.Invoke(context, nil, err)
		}
		result, err := a.useCase.Convert(criteria)
		if err != nil {
			return a.responder.Invoke(context, nil, err)
		}
		return a.responder.Invoke(context, result, err)
	}
}
