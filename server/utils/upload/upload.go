package upload

import "mime/multipart"

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	// DeleteFile(key string) error
}

func NewOss() OSS {
	return &Minio{}
}
