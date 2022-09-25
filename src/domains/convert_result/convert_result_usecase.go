package convert_result

import "log"

type UseCase interface {
	Save(content *ConvertResult) (*ConvertResult, error)
}

type useCase struct {
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &useCase{
		repository,
	}
}

func (u *useCase) Save(content *ConvertResult) (*ConvertResult, error) {
	result, err := u.repository.Save(content)
	if err != nil {
		log.Fatalf("save error:%v", err)
		return nil, err
	}

	return result, nil
}
