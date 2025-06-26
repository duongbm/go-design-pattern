package main

import "fmt"

type S3Uploader struct {
	Bucket string

	// aws client injected
}

func (u S3Uploader) Upload(file []byte, filename string) (string, error) {
	var s3Url string

	fmt.Printf("Uploading file %s to S3 Bucket: %s\n", filename, u.Bucket)

	return s3Url, nil
}
