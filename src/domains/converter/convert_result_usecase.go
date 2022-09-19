package converter

import "log"

type ConvertResultUseCase interface {
	Save(content *ConvertResult) (*ConvertResult, error)
}

type convertResultUseCase struct {
	repository ConvertResultRepository
}

func NewConvertResultUseCase(repository ConvertResultRepository) ConvertResultUseCase {
	return &convertResultUseCase{
		repository,
	}
}

func (u *convertResultUseCase) Save(content *ConvertResult) (*ConvertResult, error) {
	result, err := u.repository.Save(content)
	if err != nil {
		log.Fatalf("save error:%v", err)
		return nil, err
	}

	return result, nil
}
