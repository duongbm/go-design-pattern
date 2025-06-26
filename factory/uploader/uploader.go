package main

import (
	"fmt"
	"log"
)

type Uploader interface {
	Upload(file []byte, filename string) (string, error)
}

func UploadFactory(provider string) (Uploader, error) {
	switch provider {
	case "s3":
		return S3Uploader{}, nil
	case "gcs":
		return CloudStorageUploader{}, nil
	case "blob":
		return BlobUploader{}, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}

func main() {
	uploader, err := UploadFactory("s3")
	if err != nil {
		log.Fatal(err)
	}

	file := []byte("Hello World")
	uploadedUrl, err := uploader.Upload(file, "hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("upload success, url: ", uploadedUrl)

}
