package converter

import (
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
	"log"
)

type SpeechToTextUseCase interface {
	Convert(criteria *SpeechToTextCriteria) ([]string, error)
}

type speechToTextUseCase struct {
	translator           SpeechToTextTranslator
	repository           SpeechToTextRepository
	convertResultUseCase ConvertResultUseCase
	uploaderUseCase      uploader.UseCase
}

func NewSpeechToTextUseCase(
	convertResultUseCase ConvertResultUseCase,
	translator SpeechToTextTranslator,
	uploaderUseCase uploader.UseCase,
	repository SpeechToTextRepository,
) SpeechToTextUseCase {
	return &speechToTextUseCase{
		translator,
		repository,
		convertResultUseCase,
		uploaderUseCase,
	}
}

func (u *speechToTextUseCase) Convert(criteria *SpeechToTextCriteria) ([]string, error) {
	content := u.translator.CriteriaToUploadFile(criteria)
	fileUri, err := u.uploaderUseCase.Upload(content)
	if err != nil {
		log.Fatalf("upload error:%v", err)
		return nil, err
	}
	criteria.FileURI = fileUri
	result, err := u.repository.Convert(criteria)
	if err != nil {
		log.Fatalf("convert error:%v", err)
		return nil, err
	}

	return result, nil
}
