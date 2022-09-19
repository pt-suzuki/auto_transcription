package converter

func ProviderSpeechToTextRepository() SpeechToTextRepository {
	translator := ProviderSpeechToTextTranslator()

	return NewSpeechToTextRepository(translator)
}