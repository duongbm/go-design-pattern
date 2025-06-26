package main

import "fmt"

type BlobUploader struct {
	BlobName string
	// az client injected
}

func (b BlobUploader) Upload(file []byte, filename string) (string, error) {
	var uploadUrl string
	fmt.Printf("Uploading file %s to Blob %s\n", filename, b.BlobName)
	return uploadUrl, nil
}
