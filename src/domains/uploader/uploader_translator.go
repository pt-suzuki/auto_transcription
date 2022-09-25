package uploader

import firestore2 "cloud.google.com/go/firestore"

type Translator interface {
	ContentToMap(item *UploadFile) map[string]interface{}
	DocumentSnapshotToContent(doc *firestore2.DocumentSnapshot) *UploadFile
}

type translator struct {
}

func NewTranslator() Translator {
	return &translator{}
}

func (t *translator) ContentToMap(item *UploadFile) map[string]interface{} {
	m := make(map[string]interface{})

	m["FileName"] = item.FileName
	m["FilePath"] = item.FilePath
	m["Type"] = item.Type

	return m
}

func (t *translator) DocumentSnapshotToContent(doc *firestore2.DocumentSnapshot) *UploadFile {
	content := &UploadFile{}
	data := doc.Data()

	content.ID = doc.Ref.ID
	if data["FileName"] != nil {
		content.FileName = data["FileName"].(string)
	}
	if data["UploadFileID"] != nil {
		content.FilePath = data["FilePath"].(string)
	}
	return content
}
