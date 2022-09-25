package uploader

import (
	firestore2 "github.com/pt-suzuki/auto_transcription/infrastructure/firestore"
	"github.com/pt-suzuki/auto_transcription/src/provider/test/convert_result"
	"testing"
)

func TestUseCase_Upload(t *testing.T) {

}

func ProviderUseCase() (UseCase, error) {
	repository, err := ProviderRepository()
	if err != nil {
		return nil, err
	}

	client := firestore2.GetLocalClient()
	convertResultUseCase := convert_result.ProviderConvertResultUseCase(client)

	return NewUseCase(repository, convertResultUseCase), nil
}
