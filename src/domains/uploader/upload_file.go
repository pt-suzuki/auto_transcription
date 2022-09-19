package uploader

type UploadFile struct {
	FileName string
	Data     string
	Type     BucketType
}

type BucketType string

const (
	Speech BucketType = "speech"
)
