package echo

import "github.com/pt-suzuki/auto_transcription/src/controllers/echo/converter"

type Provider struct {
	ConvertSpeechToTextAction converter.ConvertSpeechToTextAction
}

func NewProvider(
	convertSpeechToTextAction converter.ConvertSpeechToTextAction,
) *Provider {
	return &Provider{
		convertSpeechToTextAction,
	}
}
