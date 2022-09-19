package uploader

import "testing"

func TestUseCase_Upload(t *testing.T) {

}

func ProviderUseCase() (UseCase, error) {
	repository, err := ProviderRepository()
	if err != nil {
		return nil, err
	}

	return NewUseCase(repository), nil
}
