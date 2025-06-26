package main

import "fmt"

type CloudStorageUploader struct {
	Bucket string
	// gcp client injected
}

func (g CloudStorageUploader) Upload(file []byte, filename string) (string, error) {
	var url string
	fmt.Printf("Uploading file %s to GCS Bucket: %s\n", filename, g.Bucket)
	return url, nil
}
