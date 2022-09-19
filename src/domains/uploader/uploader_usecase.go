package uploader

type UseCase interface {
	Upload(content *UploadFile) (string, error)
}

type useCase struct {
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &useCase{
		repository,
	}
}

func (u *useCase) Upload(content *UploadFile) (string, error) {
	result, err := u.repository.Upload(content)
	if err != nil {
		return "", err
	}
	return result, nil
}
