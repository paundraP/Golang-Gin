package pkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type S3aws struct {
	client *s3.Client
	bucket string
	region string
}

func NewS3AWS() *S3aws {
	bucket := os.Getenv("bucket")
	region := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		)),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to load AWS configuration: %v", err))
	}
	client := s3.NewFromConfig(cfg)

	return &S3aws{
		client: client,
		bucket: bucket,
		region: region,
	}
}

func (a *S3aws) FileUpload(fileName string, filebody *multipart.FileHeader, folderName string, mv ...string) (string, error) {
	if filebody == nil {
		return "", nil
	}

	file, err := filebody.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimetype, err := GetMimetype(file)
	if err != nil {
		return "", err
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	if len(mv) > 0 {
		valid := false
		for _, m := range mv {
			if mimetype == m {
				valid = true
				break
			}
		}

		if !valid {
			return "", fmt.Errorf("invalid mimetype: %s", mimetype)
		}
	}

	objectKey := fmt.Sprintf("%s/%s", folderName, fileName)

	_, err = a.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

func (a *S3aws) UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimetype, err := GetMimetype(file)
	if err != nil {
		return "", err
	}

	if len(mv) > 0 {
		flag := false
		for _, m := range mv {
			if mimetype == m {
				flag = true
				break
			}
		}

		if !flag {
			return "", fmt.Errorf("invalid mimetype")
		}
	}

	_, err = a.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return "", errors.New("ini error")
	}

	return objectKey, nil
}

func (a *S3aws) GetObjectKeyFromLink(link string) string {
	pref := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", a.bucket, a.region)

	if !strings.HasPrefix(link, pref) {
		return ""
	}

	objectKey := strings.TrimPrefix(link, pref)
	return objectKey
}

func (a *S3aws) GetPublicLink(objectKey string) string {
	publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", a.bucket, a.region, objectKey)
	return publicURL
}

func GetMimetype(f multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	return mimeType, nil
}
