package converter

type SpeechToTextCriteria struct {
	Data         string `json:"data"`
	FileName     string `json:"fileName"`
	FilePath     string `json:"filePath"`
	UploadFileID string `json:"uploadFileID"`
}
