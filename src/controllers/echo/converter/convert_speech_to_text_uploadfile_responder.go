package converter

import (
	"github.com/labstack/echo"
	"net/http"
)

type ConvertSpeechToTextUploadFileResponder interface {
	Invoke(context echo.Context, content []string, err error) error
}

type convertSpeechToTextUploadFileResponder struct {
}

func NewConvertSpeechToTextUploadFileResponder() ConvertSpeechToTextUploadFileResponder {
	return &convertSpeechToTextUploadFileResponder{}
}

func (*convertSpeechToTextUploadFileResponder) Invoke(context echo.Context, content []string, err error) error {
	if err != nil {
		return context.JSON(http.StatusBadRequest, []interface{}{"bad request"})
	}
	return context.JSON(http.StatusOK, content)
}
