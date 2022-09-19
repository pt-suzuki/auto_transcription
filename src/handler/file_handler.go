package handler

import (
	"encoding/base64"
	"log"
	"os"
)

func FileEncode(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("fail file open : %v", err)
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	fi, err := file.Stat() //FileInfo interface
	if err != nil {
		log.Fatalf("fail stat : %v", err)
		return "", err
	}
	size := fi.Size() //ファイルサイズ

	data := make([]byte, size)
	_, _ = file.Read(data)

	return base64.StdEncoding.EncodeToString(data), nil
}
