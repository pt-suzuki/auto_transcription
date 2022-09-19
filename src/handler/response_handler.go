package handler

import (
	"bytes"
	"github.com/labstack/echo"
	"io/ioutil"
)

type ResponseHandler interface {
	GetBodyByResponse(context echo.Context) ([]byte, error)
}

type responseHandler struct {
}

func NewResponseHandler() ResponseHandler {
	return &responseHandler{}
}

func (rh *responseHandler) GetBodyByResponse(context echo.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(context.Request().Body)
	context.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, err
}
