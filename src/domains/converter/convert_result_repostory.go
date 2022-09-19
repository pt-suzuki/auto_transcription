package converter

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

const collectionName = "convert_result"

type ConvertResultRepository interface {
	Save(item *ConvertResult) (*ConvertResult, error)
	GetListByCriteria(_ *ConvertResultSearchCriteria) []*ConvertResult
	GetContentById(id string) (*ConvertResult, error)
}

type convertResultRepository struct {
	client     *firestore.Client
	translator ConvertResultTranslator
}

func NewConvertResultRepository(client *firestore.Client, translator ConvertResultTranslator) ConvertResultRepository {
	return &convertResultRepository{
		client,
		translator,
	}
}

func (r *convertResultRepository) Save(item *ConvertResult) (*ConvertResult, error) {
	m := r.translator.ContentToMap(item)

	ref, _, err := r.client.Collection(collectionName).Add(context.Background(), m)
	if err != nil {
		log.Fatalf("fail collection  %s save: %v", collectionName, err)
		return nil, err
	}
	item.ID = ref.ID

	return item, nil
}

func (r *convertResultRepository) GetListByCriteria(_ *ConvertResultSearchCriteria) []*ConvertResult {
	ctx := context.Background()
	iterator := r.client.Collection(collectionName).Documents(ctx)
	return r.translator.IteratorToList(iterator)
}

func (r *convertResultRepository) GetContentById(id string) (*ConvertResult, error) {
	ctx := context.Background()
	snap, err := r.client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	return r.translator.DocumentSnapshotToContent(snap), nil
}
