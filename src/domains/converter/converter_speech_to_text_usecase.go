package converter

import (
	"github.com/pt-suzuki/auto_transcription/src/domains/convert_result"
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
	"log"
	"time"
)

type SpeechToTextUseCase interface {
	UploadAndConvert(criteria *SpeechToTextCriteria) ([]string, error)
	ConvertByUploadFileId(id string) ([]string, error)
	Convert(criteria *SpeechToTextCriteria) ([]string, error)
}

type speechToTextUseCase struct {
	translator           SpeechToTextTranslator
	repository           SpeechToTextRepository
	convertResultUseCase convert_result.UseCase
	uploaderUseCase      uploader.UseCase
}

func NewSpeechToTextUseCase(
	convertResultUseCase convert_result.UseCase,
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

func (u *speechToTextUseCase) UploadAndConvert(criteria *SpeechToTextCriteria) ([]string, error) {
	content := u.translator.CriteriaToUploadFile(criteria)
	uploadFile, err := u.uploaderUseCase.Upload(content)
	if err != nil {
		return nil, err
	}
	criteria.FilePath = uploadFile.FilePath
	criteria.UploadFileID = uploadFile.ID

	result, err := u.Convert(criteria)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *speechToTextUseCase) ConvertByUploadFileId(id string) ([]string, error) {
	uploadFile, err := u.uploaderUseCase.Get(id)
	if err != nil {
		return nil, err
	}
	criteria := &SpeechToTextCriteria{
		FilePath: uploadFile.FilePath,
		FileName: uploadFile.FileName,
	}
	result, err := u.Convert(criteria)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *speechToTextUseCase) Convert(criteria *SpeechToTextCriteria) ([]string, error) {
	result, err := u.repository.Convert(criteria)
	if err != nil {
		log.Fatalf("convert error:%v", err)
		return nil, err
	}
	convertResult := &convert_result.ConvertResult{
		UploadFileID:  criteria.UploadFileID,
		ConvertResult: result,
		CreatedAt:     time.Now(),
	}
	if _, err = u.convertResultUseCase.Save(convertResult); err != nil {
		return nil, err
	}
	return result, nil
}
