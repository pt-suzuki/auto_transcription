package router

import (
	echo2 "github.com/labstack/echo"

	"github.com/pt-suzuki/auto_transcription/src/provider/echo"
)

func GetRouter(provider *echo.Provider) *echo2.Echo {
	e := echo2.New()

	// ftvMiddleware := provider.MiddlewareProvider.FirebaseTokenVerifiedMiddleware.GetFirebaseTokenVerifiedMiddleware()

	api := e.Group("/api")
	api.POST("/converter/speech",
		provider.ControllerProvider.ConverterControllerProvider.ConvertSpeechToTextUploadFileAction.Invoke(),
		//ftvMiddleware)
	)
	api.POST("/converter/speech/:upload_file_id",
		provider.ControllerProvider.ConverterControllerProvider.ConvertSpeechToTextByUploadFileIDAction.Invoke(),
		//ftvMiddleware)
	)

	return e
}
