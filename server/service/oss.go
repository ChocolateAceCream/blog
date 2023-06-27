package service

import (
	"mime/multipart"

	"github.com/ChocolateAceCream/blog/utils/upload"
)

type OssService struct{}

func (os *OssService) UploadFile(header *multipart.FileHeader) (path string, err error) {
	oss := upload.NewOss()
	path, err = oss.UploadFile(header)
	if err != nil {
		panic(err)
	}
	return path, err
}
