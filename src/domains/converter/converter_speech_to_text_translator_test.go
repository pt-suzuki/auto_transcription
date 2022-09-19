package converter

import "github.com/pt-suzuki/auto_transcription/src/provider/test/server"

func ProviderSpeechToTextTranslator() SpeechToTextTranslator {
	handler := server.ProviderResponseHandler()
	return NewSpeechToTextTranslator(handler)
}
