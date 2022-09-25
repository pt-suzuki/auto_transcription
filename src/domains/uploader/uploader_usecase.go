package uploader

import (
	"github.com/pt-suzuki/auto_transcription/src/domains/convert_result"
	"log"
)

type UseCase interface {
	Upload(content *UploadFile) (*UploadFile, error)
	Save(content *UploadFile) (*UploadFile, error)
	Get(id string) (*UploadFile, error)
}

type useCase struct {
	repository           Repository
	convertResultUseCase convert_result.UseCase
}

func NewUseCase(repository Repository, convertResultUseCase convert_result.UseCase) UseCase {
	return &useCase{
		repository,
		convertResultUseCase,
	}
}

func (u *useCase) Upload(content *UploadFile) (*UploadFile, error) {
	path, err := u.repository.Upload(content)
	if err != nil {
		log.Fatalf("upload error:%v", err)
		return nil, err
	}

	content.FilePath = path
	result, err := u.Save(content)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *useCase) Save(content *UploadFile) (*UploadFile, error) {
	result, err := u.repository.Save(content)
	if err != nil {
		log.Fatalf("upload file save error:%v", err)
		return nil, err
	}
	return result, nil
}

func (u *useCase) Get(id string) (*UploadFile, error) {
	result, err := u.repository.Get(id)
	if err != nil {
		log.Fatalf("upload file get error:%v", err)
		return nil, err
	}
	return result, nil
}
