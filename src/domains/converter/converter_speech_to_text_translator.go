package converter

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/pt-suzuki/auto_transcription/src/domains/uploader"
	"github.com/pt-suzuki/auto_transcription/src/handler"
	"google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type SpeechToTextTranslator interface {
	EchoContextToCriteria(context echo.Context) (*SpeechToTextCriteria, error)
	RecognizeResponseToClauseList(resp *speech.RecognizeResponse) []string
	CriteriaToUploadFile(criteria *SpeechToTextCriteria) *uploader.UploadFile
	EchoContextToId(context echo.Context) string
}

type speechToTextTranslator struct {
	responseHandler handler.ResponseHandler
}

func NewSpeechToTextTranslator(responseHandler handler.ResponseHandler) SpeechToTextTranslator {
	return &speechToTextTranslator{
		responseHandler,
	}
}

func (t *speechToTextTranslator) EchoContextToCriteria(context echo.Context) (*SpeechToTextCriteria, error) {
	result := &SpeechToTextCriteria{}
	body, err := t.responseHandler.GetBodyByResponse(context)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*speechToTextTranslator) CriteriaToUploadFile(criteria *SpeechToTextCriteria) *uploader.UploadFile {
	return &uploader.UploadFile{
		Data:     criteria.Data,
		FileName: criteria.FileName,
		Type:     uploader.Speech,
	}
}

func (*speechToTextTranslator) RecognizeResponseToClauseList(resp *speech.RecognizeResponse) []string {
	result := make([]string, 0)
	// Prints the results
	for _, item := range resp.Results {
		for _, alt := range item.Alternatives {
			result = append(result, alt.Transcript)
		}
	}
	return result
}

func (*speechToTextTranslator) EchoContextToId(context echo.Context) string {
	return context.Param("upload_file_id")
}
