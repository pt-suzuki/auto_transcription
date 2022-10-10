package converter

import (
	"github.com/labstack/echo"
	"net/http"
)

type ConvertSpeechToTextByUploadFileIDResponder interface {
	Invoke(context echo.Context, content []string, err error) error
}

type convertSpeechToTextByUploadFileIDResponder struct {
}

func NewConvertSpeechToTextByUploadFileIDResponder() ConvertSpeechToTextByUploadFileIDResponder {
	return &convertSpeechToTextByUploadFileIDResponder{}
}

func (*convertSpeechToTextByUploadFileIDResponder) Invoke(context echo.Context, content []string, err error) error {
	if err != nil {
		return context.JSON(http.StatusBadRequest, []interface{}{"bad request:" + err.Error()})
	}
	return context.JSON(http.StatusOK, content)
}
