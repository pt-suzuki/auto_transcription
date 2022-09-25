package uploader

import (
	firestore2 "github.com/pt-suzuki/auto_transcription/infrastructure/firestore"
	"github.com/pt-suzuki/auto_transcription/src/provider/test"
	"github.com/pt-suzuki/auto_transcription/src/provider/test/convert_result"
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
	fireStorageClient := firestorage.GetLocalClient()
	fireStoreClient := firestore2.GetLocalClient()

	return uploader.NewRepository(fireStoreClient, fireStorageClient, translator), nil
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

	fireStoreClient := firestore2.GetLocalClient()
	convertResultUseCase := convert_result.ProviderConvertResultUseCase(fireStoreClient)

	return uploader.NewUseCase(repository, convertResultUseCase), nil
}
