package converter

import (
	firestore2 "cloud.google.com/go/firestore"
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
	"github.com/pt-suzuki/auto_transcription/src/handler"
	uploader2 "github.com/pt-suzuki/auto_transcription/src/provider/test/uploader"
	"testing"
)

func TestSpeechToTextUseCase_Convert(t *testing.T) {
	/*
		useCase, err := ProviderSpeechToTextUseCase()
		if err != nil {
			t.Fatal(err)
		}

		t.Run("正常系", func(t *testing.T) {
			content, err := createUploadFile()
			if err != nil {
				t.Fatal(err)
			}

			result, err := useCase.Convert(content)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, len(result), 2)
		})
	*/
}

func ProviderSpeechToTextUseCase(client *firestore2.Client) (SpeechToTextUseCase, error) {
	uploadUseCase, err := uploader2.ProvideUploaderUseCase()
	if err != nil {
		return nil, err
	}
	translator := ProviderSpeechToTextTranslator()
	repository := ProviderSpeechToTextRepository()

	convertResultUseCase := ProviderConvertResultUseCase(client)

	return NewSpeechToTextUseCase(convertResultUseCase, translator, uploadUseCase, repository), nil
}

func createUploadFile() (*uploader.UploadFile, error) {
	data, err := handler.FileEncode("../../../test/mp3/tomato.mp3")
	if err != nil {
		return nil, err
	}

	return &uploader.UploadFile{
		Type:     uploader.Speech,
		Data:     data,
		FileName: "tomato.mp3",
	}, nil
}
