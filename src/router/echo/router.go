package echo

import (
	echo2 "github.com/labstack/echo"

	"github.com/pt-suzuki/auto_transcription/src/provider/echo"
)

func GetRouter(provider *echo.Provider) *echo2.Echo {
	e := echo2.New()

	api := e.Group("/api")
	api.POST("/converter/speech", provider.ConvertSpeechToTextAction.Invoke())

	return e
}
