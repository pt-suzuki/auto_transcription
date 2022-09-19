package converter

import (
	firestore2 "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
)

type ConvertResultTranslator interface {
	ContentToMap(item *ConvertResult) map[string]interface{}
	IteratorToList(ite *firestore2.DocumentIterator) []*ConvertResult
	DocumentSnapshotToContent(doc *firestore2.DocumentSnapshot) *ConvertResult
}

type convertResultTranslator struct {
}

func NewConvertResultTranslator() ConvertResultTranslator {
	return &convertResultTranslator{}
}

func (t *convertResultTranslator) ContentToMap(item *ConvertResult) map[string]interface{} {
	m := make(map[string]interface{})

	m["FilePath"] = item.FilePath
	m["ConvertResult"] = item.ConvertResult

	return m
}

func (t *convertResultTranslator) IteratorToList(ite *firestore2.DocumentIterator) []*ConvertResult {
	result := make([]*ConvertResult, 0)
	for {
		doc, err := ite.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		f := t.DocumentSnapshotToContent(doc)
		result = append(result, f)
	}
	return result
}

func (t *convertResultTranslator) DocumentSnapshotToContent(doc *firestore2.DocumentSnapshot) *ConvertResult {
	var content = &ConvertResult{}
	data := doc.Data()

	content.ID = doc.Ref.ID
	if data["ConvertResult"] != nil {
		for _, item := range data["ConvertResult"].([]interface{}) {
			content.ConvertResult = append(content.ConvertResult, item.(string))
		}
	}
	if data["FilePath"] != nil {
		content.FilePath = data["FilePath"].(string)
	}
	return content
}
