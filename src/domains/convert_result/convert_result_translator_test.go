package convert_result

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslator_ContentToMap(t *testing.T) {
	tl := ProviderConvertResultTranslator()
	t.Run("値オブジェクトをマップに変換", func(t *testing.T) {
		item := createConvertResult()
		result := tl.ContentToMap(item)

		t.Run("IDが空", func(t *testing.T) {
			assert.Empty(t, result["ID"])
		})
		t.Run("変換結果が入力される", func(t *testing.T) {
			t.Run("件数が2件", func(t *testing.T) {
				assert.Equal(t, len(result["ConvertResult"].([]string)), 2)
			})
			t.Run("1件目の値が一致", func(t *testing.T) {
				assert.Equal(t, result["ConvertResult"].([]string)[0], "テスト1")
			})
			t.Run("2件目の値が一致", func(t *testing.T) {
				assert.Equal(t, result["ConvertResult"].([]string)[1], "テスト2")
			})
		})
		t.Run("アップロードファイルが一致", func(t *testing.T) {
			assert.Equal(t, result["UploadFileID"], "テストアップロードファイルID")
		})
	})
}

func ProviderConvertResultTranslator() Translator {
	return NewConvertResultTranslator()
}
