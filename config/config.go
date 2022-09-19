package config

import (
	"os"
)

type Config struct {
	GcpProjectId  string
	StorageBucket string
	StorageRoot   string
	BucketPath    BucketPath
}

type BucketPath struct {
	SpeechToText string
}

func NewConfig() Config {
	conf := Config{
		GcpProjectId:  os.Getenv("GCP_PROJECT"),
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
		StorageRoot:   os.Getenv("STORAGE_ROOT"),
		BucketPath: BucketPath{
			SpeechToText: os.Getenv("STORAGE_BUCKET_SPEECH_TO_TEXT"),
		},
	}
	return conf
}
