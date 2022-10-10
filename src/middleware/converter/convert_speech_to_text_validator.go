package converter

import (
	"encoding/base64"
	"github.com/labstack/echo"

	"github.com/pt-suzuki/auto_transcription/src/domains/converter"
	"net/http"
	"strings"
)

type ConvertSpeechToTextValidatorMiddleware interface {
	GetCreateValidatorMiddleware() echo.MiddlewareFunc
}

type convertSpeechToTextValidatorMiddleware struct {
	translator converter.SpeechToTextTranslator
}

func NewConvertSpeechToTextValidatorMiddleware(translator converter.SpeechToTextTranslator) ConvertSpeechToTextValidatorMiddleware {
	return &convertSpeechToTextValidatorMiddleware{translator}
}

func (m *convertSpeechToTextValidatorMiddleware) GetCreateValidatorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			content, err := m.translator.EchoContextToCriteria(c)
			if err != nil {
				return echo.NewHTTPError(400, "Bad Request")
			}
			if content.FileName == "" {
				return echo.NewHTTPError(422, "Required FileName")
			}
			if content.Data == "" {
				return echo.NewHTTPError(422, "Required Data")
			}

			decoded, err := base64.StdEncoding.DecodeString(content.Data)
			if err != nil {
				return echo.NewHTTPError(400, "File Decode Fail")
			}
			mimeType := http.DetectContentType(decoded)
			if !strings.Contains(mimeType, "audio") && !strings.Contains(mimeType, "video") {
				return echo.NewHTTPError(422, "Not Allow MimeType")
			}

			err = next(c)
			return err
		}
	}
}
