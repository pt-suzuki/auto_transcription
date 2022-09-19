package converter

import "time"

type ConvertResult struct {
	ID            string
	FilePath      string
	ConvertResult []string
	CreatedAt     time.Time
}
