package storageSdk

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService interface {
	UploadFile(bucketName string, file *multipart.FileHeader) (string, error)
	DeleteFile(bucketName string, fileName string) error
	CreateBucket(bucketName string) error
}

type minioService struct {
	minioClient *minio.Client
}

func NewMinioService(
	host string,
	user string,
	password string,
	useSSL bool,
) (MinioService, error) {
	minioClient, err := minio.New(host, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &minioService{
		minioClient: minioClient,
	}, nil
}

func (s *minioService) UploadFile(bucketName string, file *multipart.FileHeader) (string, error) {

	if file == nil {
		return "", errors.New("Файл хоосон байна")
	}

	bucketExists, err := s.IsBucketExists(bucketName)
	if err != nil {
		return "", err
	}

	if !bucketExists {
		err = s.CreateBucket(bucketName)
		if err != nil {
			return "", err
		}
	}

	openFile, err := file.Open()
	if err != nil {
		return "", err
	}

	types := file.Header["Content-Type"]

	info, err := s.minioClient.PutObject(context.Background(), bucketName, file.Filename, openFile, file.Size, minio.PutObjectOptions{
		ContentType:  types[0],
		CacheControl: "max-age=31536000, immutable",
	})
	if err != nil {
		return "", err
	}

	return bucketName + "/" + info.Key, nil
}

func (s *minioService) IsBucketExists(bucketName string) (bool, error) {

	exists, err := s.minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *minioService) DeleteFile(bucketName string, fileName string) error {

	err := s.minioClient.RemoveObject(context.Background(), bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *minioService) CreateBucket(bucketName string) error {

	err := s.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}

	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject","s3:PutObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + bucketName + `/*"],"Sid": ""}]}`

	if err = s.minioClient.SetBucketPolicy(context.Background(), bucketName, policy); err != nil {
		return err
	}

	return nil
}
