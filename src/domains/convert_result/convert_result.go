package convert_result

import "time"

type ConvertResult struct {
	ID           string
	UploadFileID string
	Results      []string `json:"results"`
	CreatedAt    time.Time
}
