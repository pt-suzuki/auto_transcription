package uploader

import (
	"bytes"
	"cloud.google.com/go/firestore"
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

const collectionName = "upload_file"

type Repository interface {
	Upload(data *UploadFile) (string, error)
	Save(item *UploadFile) (*UploadFile, error)
	Get(id string) (*UploadFile, error)
}

type repository struct {
	firestoreClient *firestore.Client
	storageClient   *storage.Client
	translator      Translator
}

func NewRepository(firestoreClient *firestore.Client, storageClient *storage.Client, translator Translator) Repository {
	return &repository{
		firestoreClient,
		storageClient,
		translator,
	}
}

func (r *repository) Upload(content *UploadFile) (string, error) {
	path := r.createPath(content)
	conf := config.NewConfig()
	bucket, err := r.storageClient.Bucket(conf.StorageBucket)
	if err != nil {
		log.Printf("get default bucket error:%v", err)
		return "", err
	}
	decoded, err := base64.StdEncoding.DecodeString(content.Data)
	if err != nil {
		log.Printf("decode error:%v", err)
		return "", err
	}
	object := bucket.Object(path)

	ctx := context.Background()
	writer := object.NewWriter(ctx)
	if _, err = io.Copy(writer, bytes.NewReader(decoded)); err != nil {
		log.Printf("file copy error:%v", err)
		return "", err
	}
	if err := writer.Close(); err != nil {
		log.Printf("file create error:%v", err)
		return "", err
	}
	if err := object.ACL().Set(context.Background(), storage2.AllUsers, storage2.RoleReader); err != nil {
		log.Printf("set acl error:%v", err)
		return "", err
	}
	return path, nil
}

func (r *repository) Save(item *UploadFile) (*UploadFile, error) {
	m := r.translator.ContentToMap(item)

	ref, _, err := r.firestoreClient.Collection(collectionName).Add(context.Background(), m)
	if err != nil {
		log.Printf("fail collection  %s save: %v", collectionName, err)
		return nil, err
	}
	item.ID = ref.ID

	return item, nil
}

func (r *repository) Get(id string) (*UploadFile, error) {
	ctx := context.Background()
	snap, err := r.firestoreClient.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	return r.translator.DocumentSnapshotToContent(snap), nil
}

func (r *repository) createPath(content *UploadFile) string {
	conf := config.NewConfig()
	date := time.Now()
	strMonth := date.Format("200601")
	strDate := date.Format("20060102")
	path := fmt.Sprintf("%s/%s/%s/%s/%s", conf.StorageRoot, content.Type, strMonth, strDate, content.FileName)

	return path
}
