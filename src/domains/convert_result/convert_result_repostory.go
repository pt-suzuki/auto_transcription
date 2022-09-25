package convert_result

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

const collectionName = "convert_result"

type Repository interface {
	Save(item *ConvertResult) (*ConvertResult, error)
	GetListByCriteria(_ *SearchCriteria) []*ConvertResult
	GetContentById(id string) (*ConvertResult, error)
}

type repository struct {
	client     *firestore.Client
	translator Translator
}

func NewRepository(client *firestore.Client, translator Translator) Repository {
	return &repository{
		client,
		translator,
	}
}

func (r *repository) Save(item *ConvertResult) (*ConvertResult, error) {
	m := r.translator.ContentToMap(item)

	ref, _, err := r.client.Collection(collectionName).Add(context.Background(), m)
	if err != nil {
		log.Fatalf("fail collection %s save: %v", collectionName, err)
		return nil, err
	}
	item.ID = ref.ID

	return item, nil
}

func (r *repository) GetListByCriteria(_ *SearchCriteria) []*ConvertResult {
	ctx := context.Background()
	iterator := r.client.Collection(collectionName).Documents(ctx)
	return r.translator.IteratorToList(iterator)
}

func (r *repository) GetContentById(id string) (*ConvertResult, error) {
	ctx := context.Background()
	snap, err := r.client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	return r.translator.DocumentSnapshotToContent(snap), nil
}
