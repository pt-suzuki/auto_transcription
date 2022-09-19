package uploader

import (
	"github.com/pt-suzuki/auto_transcription/src/handler"
	"github.com/pt-suzuki/auto_transcription/src/provider/test"
	"log"

	"github.com/pt-suzuki/auto_transcription/infrastructure/firestorage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Upload(t *testing.T) {
	r, err := ProviderRepository()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("正常系", func(t *testing.T) {
		content, err := createUploadFile()
		if err != nil {
			t.Fatal(err)
		}
		t.Run("ファイルパスが返る", func(t *testing.T) {
			filePath, err := r.Upload(content)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, filePath, "auto_transcription/speech/sea_test2.jpg")
		})
	})
}

func createUploadFile() (*UploadFile, error) {
	data, err := handler.FileEncode("../../../test/img/sea.jpg")
	if err != nil {
		return nil, err
	}

	return &UploadFile{
		Type:     Speech,
		Data:     data,
		FileName: "sea_test2.jpg",
	}, nil
}

func ProviderRepository() (Repository, error) {
	err := test.Init()
	if err != nil {
		log.Fatalf("env not read: %v", err)
		return nil, err
	}
	translator := ProviderTranslator()
	client := firestorage.GetLocalClient()
	return NewRepository(client, translator), nil
}
