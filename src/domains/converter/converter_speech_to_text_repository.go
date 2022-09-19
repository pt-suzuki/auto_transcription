package converter

import (
	"context"
	"fmt"
	"github.com/pt-suzuki/auto_transcription/config"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
)

type SpeechToTextRepository interface {
	Convert(criteria *SpeechToTextCriteria) ([]string, error)
}

type speechToTextRepository struct {
	translator SpeechToTextTranslator
}

func NewSpeechToTextRepository(translator SpeechToTextTranslator) SpeechToTextRepository {
	return &speechToTextRepository{
		translator,
	}
}

func (r *speechToTextRepository) Convert(criteria *SpeechToTextCriteria) ([]string, error) {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	defer func() {
		_ = client.Close()
	}()

	conf := config.NewConfig()
	path := fmt.Sprintf("gs://%s/%s", conf.StorageBucket, criteria.FileURI)

	// Detects speech in the audio file
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			SampleRateHertz: 44100,
			LanguageCode:    "ja-JP",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: path},
		},
	})

	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
		return nil, err
	}

	return r.translator.RecognizeResponseToClauseList(resp), nil
}
