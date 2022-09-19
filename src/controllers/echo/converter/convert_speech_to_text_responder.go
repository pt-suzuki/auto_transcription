package converter

import (
	"github.com/labstack/echo"
	"net/http"
)

type ConvertSpeechToTextResponder interface {
	Invoke(context echo.Context, content []string, err error) error
}

type convertSpeechToTextResponder struct {
}

func NewConvertSpeechToTextResponder() ConvertSpeechToTextResponder {
	return &convertSpeechToTextResponder{}
}

func (*convertSpeechToTextResponder) Invoke(context echo.Context, content []string, err error) error {
	if err != nil {
		return context.JSON(http.StatusBadRequest, []interface{}{"bad request"})
	}
	return context.JSON(http.StatusOK, content)
}
