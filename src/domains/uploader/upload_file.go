package uploader

type UploadFile struct {
	ID       string
	FileName string
	FilePath string
	Data     string
	Type     BucketType
}

type BucketType string

const (
	Speech BucketType = "speech"
)
