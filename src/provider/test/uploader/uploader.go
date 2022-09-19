package uploader

import (
	"github.com/pt-suzuki/auto_transcription/src/provider/test"
	"log"

	"github.com/pt-suzuki/auto_transcription/infrastructure/firestorage"
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
)

func ProvideUploaderTranslator() (uploader.Translator, error) {
	err := test.Init()
	if err != nil {
		log.Fatalf("env not read: %v", err)
		return nil, err
	}
	return uploader.NewTranslator(), nil
}

func ProvideUploaderRepository() (uploader.Repository, error) {
	err := test.Init()
	if err != nil {
		log.Fatalf("env not read: %v", err)
		return nil, err
	}
	translator, err := ProvideUploaderTranslator()
	if err != nil {
		log.Fatalf("fail provide uploader translator: %v", err)
		return nil, err
	}
	client := firestorage.GetLocalClient()

	return uploader.NewRepository(client, translator), nil
}

func ProvideUploaderUseCase() (uploader.UseCase, error) {
	err := test.Init()
	if err != nil {
		log.Fatalf("env not read: %v", err)
		return nil, err
	}
	repository, err := ProvideUploaderRepository()
	if err != nil {
		log.Fatalf("fail provide uploader repository: %v", err)
		return nil, err
	}

	return uploader.NewUseCase(repository), nil
}
