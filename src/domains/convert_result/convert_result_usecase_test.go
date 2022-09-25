package convert_result

import firestore2 "cloud.google.com/go/firestore"

func ProviderConvertResultUseCase(client *firestore2.Client) UseCase {
	repository := ProviderConvertResultRepository(client)

	return NewUseCase(repository)
}
