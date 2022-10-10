package convert_result

import (
	firestore2 "cloud.google.com/go/firestore"
	"github.com/pt-suzuki/auto_transcription/infrastructure/firestore"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository_Save(t *testing.T) {
	client := firestore.GetLocalClient()
	r := ProviderConvertResultRepository(client)

	defer func() {
		_ = client.Close()
	}()

	t.Run("正常系", func(t *testing.T) {
		item := createConvertResult()
		result, err := r.Save(item)
		if err != nil {
			t.Fatal(err)
		}

		t.Run("IDが入力される", func(t *testing.T) {
			assert.NotEqual(t, result.ID, "テストID")
		})

		content, err := r.GetContentById(result.ID)
		if err != nil {
			t.Fatal(err)
		}
		t.Run("変換結果が一致", func(t *testing.T) {
			t.Run("件数が2件", func(t *testing.T) {
				assert.Equal(t, len(content.Results), 2)
			})
			t.Run("1件目の値が一致", func(t *testing.T) {
				assert.Equal(t, content.Results[0], "テスト1")
			})
			t.Run("2件目の値が一致", func(t *testing.T) {
				assert.Equal(t, content.Results[1], "テスト2")
			})
		})
		t.Run("アップロードファイルIDが一致", func(t *testing.T) {
			assert.Equal(t, content.UploadFileID, "テストアップロードファイルID")
		})
	})
}

func createConvertResult() *ConvertResult {
	return &ConvertResult{
		ID:           "テストID",
		UploadFileID: "テストアップロードファイルID",
		Results: []string{
			"テスト1",
			"テスト2",
		},
		CreatedAt: time.Now(),
	}
}

func ProviderConvertResultRepository(client *firestore2.Client) Repository {
	translator := ProviderConvertResultTranslator()

	return NewRepository(client, translator)
}
