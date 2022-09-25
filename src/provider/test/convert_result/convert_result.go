package convert_result

import (
	firestore2 "cloud.google.com/go/firestore"
	"github.com/pt-suzuki/auto_transcription/src/domains/convert_result"
)

func ProviderConvertResultRepository(client *firestore2.Client) convert_result.Repository {

	translator := ProviderConvertResultTranslator()

	return convert_result.NewRepository(client, translator)
}

func ProviderConvertResultTranslator() convert_result.Translator {
	return convert_result.NewConvertResultTranslator()
}

func ProviderConvertResultUseCase(client *firestore2.Client) convert_result.UseCase {
	repository := ProviderConvertResultRepository(client)

	return convert_result.NewUseCase(repository)
}
