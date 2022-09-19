package converter

import (
	firestore2 "cloud.google.com/go/firestore"
	"github.com/pt-suzuki/auto_transcription/infrastructure/firestore"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvertResultRepository_Save(t *testing.T) {
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
				assert.Equal(t, len(content.ConvertResult), 2)
			})
			t.Run("1件目の値が一致", func(t *testing.T) {
				assert.Equal(t, content.ConvertResult[0], "テスト1")
			})
			t.Run("2件目の値が一致", func(t *testing.T) {
				assert.Equal(t, content.ConvertResult[1], "テスト2")
			})
		})
		t.Run("ファイルパスが一致", func(t *testing.T) {
			assert.Equal(t, content.FilePath, "テストファイルパス")
		})
	})
}

func createConvertResult() *ConvertResult {
	return &ConvertResult{
		ID:       "テストID",
		FilePath: "テストファイルパス",
		ConvertResult: []string{
			"テスト1",
			"テスト2",
		},
		CreatedAt: time.Now(),
	}
}

func ProviderConvertResultRepository(client *firestore2.Client) ConvertResultRepository {

	translator := ProviderConvertResultTranslator()

	return NewConvertResultRepository(client, translator)
}
