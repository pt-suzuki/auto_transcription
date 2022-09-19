package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileEncode(t *testing.T) {
	t.Run("ファイルをbase64にエンコード", func(t *testing.T) {
		data, err := FileEncode("../../test/img/sea.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, data)
	})
}
