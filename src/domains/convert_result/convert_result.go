package convert_result

import "time"

type ConvertResult struct {
	ID            string
	UploadFileID  string
	ConvertResult []string
	CreatedAt     time.Time
}
