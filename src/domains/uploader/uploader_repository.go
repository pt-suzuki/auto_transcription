package uploader

import (
	"bytes"
	storage2 "cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"firebase.google.com/go/storage"
	"fmt"
	"github.com/pt-suzuki/auto_transcription/config"
	"io"
	"log"
	"time"
)

type Repository interface {
	Upload(data *UploadFile) (string, error)
}

type repository struct {
	client     *storage.Client
	translator Translator
}

func NewRepository(client *storage.Client, translator Translator) Repository {
	return &repository{
		client,
		translator,
	}
}

func (r *repository) Upload(content *UploadFile) (string, error) {
	bucket, err := r.client.DefaultBucket()
	if err != nil {
		log.Fatalf("get default bucket error:%v", err)
		return "", err
	}
	decoded, err := base64.StdEncoding.DecodeString(content.Data)
	if err != nil {
		log.Fatalf("decode error:%v", err)
		return "", err
	}

	conf := config.NewConfig()
	ctx := context.Background()
	date := time.Now()
	strMonth := date.Format("200601")
	strDate := date.Format("20060102")
	path := fmt.Sprintf("%s/%s/%s/%s/%s", conf.StorageRoot, content.Type, strMonth, strDate, content.FileName)

	object := bucket.Object(path)
	writer := object.NewWriter(ctx)

	if _, err = io.Copy(writer, bytes.NewReader(decoded)); err != nil {
		log.Fatalf("file copy error:%v", err)
		return "", err
	}

	if err := writer.Close(); err != nil {
		log.Fatalf("file create error:%v", err)
		return "", err
	}

	if err := object.ACL().Set(context.Background(), storage2.AllUsers, storage2.RoleReader); err != nil {
		log.Fatalf("set acl error:%v", err)
		return "", err
	}

	return path, nil
}
