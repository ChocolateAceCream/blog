package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type Minio struct{}

func newClient() *minio.Client {
	config := global.CONFIG.Minio
	minioClient, err := minio.New(config.Host+":"+config.Port, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Username, config.Password, ""),
		Secure: config.Https,
	})
	if err != nil {
		global.LOGGER.Error("fail to init minio client", zap.Error(err))
	}

	// Make a new bucket.
	bucketName := config.Bucket
	region := config.Region

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			global.LOGGER.Info(fmt.Sprintf("We already own %s\n", bucketName))
			return minioClient
		} else {
			global.LOGGER.Error("fail to find bucket", zap.Error(err))
			return nil
		}
	} else {
		return minioClient
	}
}

func (*Minio) UploadFile(file *multipart.FileHeader) (string, error) {
	config := global.CONFIG.Minio
	minioClient := newClient()
	filename := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	f, openError := file.Open()
	if openError != nil {
		global.LOGGER.Error("function file.Open() Filed", zap.Error(openError))
		return "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	opts := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	uploadInfo, err := minioClient.PutObject(context.Background(), config.Bucket, filename, f, file.Size, opts)
	if err != nil {
		global.LOGGER.Error("function uploader.Upload() Filed", zap.Error(err))
		return "", err
	}
	url := "https://" + config.Host + "/minio/" + config.Bucket + "/" + uploadInfo.Key
	return url, nil
}
