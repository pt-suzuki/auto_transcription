package converter

import firestore2 "cloud.google.com/go/firestore"

func ProviderConvertResultUseCase(client *firestore2.Client) ConvertResultUseCase {
	repository := ProviderConvertResultRepository(client)

	return NewConvertResultUseCase(repository)
}
