package converter

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/auto_transcription/src/domains/convert_result"
	"net/http"
)

type ConvertSpeechToTextUploadFileResponder interface {
	Invoke(context echo.Context, content *convert_result.ConvertResult, err error) error
}

type convertSpeechToTextUploadFileResponder struct {
}

func NewConvertSpeechToTextUploadFileResponder() ConvertSpeechToTextUploadFileResponder {
	return &convertSpeechToTextUploadFileResponder{}
}

func (*convertSpeechToTextUploadFileResponder) Invoke(context echo.Context, content *convert_result.ConvertResult, err error) error {
	if err != nil {
		return context.JSON(http.StatusBadRequest, []interface{}{"Internal Error"})
	}
	return context.JSON(http.StatusOK, content)
}
