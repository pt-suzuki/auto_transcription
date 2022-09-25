package controllers

import (
	converter2 "github.com/pt-suzuki/auto_transcription/src/controllers/echo/converter"
)

type ConverterControllerProvider struct {
	ConvertSpeechToTextUploadFileAction     converter2.ConvertSpeechToTextUploadFileAction
	ConvertSpeechToTextByUploadFileIDAction converter2.ConvertSpeechToTextByUploadFileIDAction
}

func NewConverterMiddlewareProvider(
	ConvertSpeechToTextUploadFileAction converter2.ConvertSpeechToTextUploadFileAction,
	ConvertSpeechToTextByUploadFileIDAction converter2.ConvertSpeechToTextByUploadFileIDAction,
) *ConverterControllerProvider {
	return &ConverterControllerProvider{
		ConvertSpeechToTextUploadFileAction,
		ConvertSpeechToTextByUploadFileIDAction,
	}
}
