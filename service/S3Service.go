package service

import (
	"bytes"
	"economic-school/util"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Uploader struct {
	svc *s3.S3
}

func NewS3Uploader(config util.Config) (*S3Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AwsRegion),
		Credentials: credentials.NewStaticCredentials(
			config.AWSAccessKey,
			config.AWSSecretAccessKey,
			""), // 空のトークンを指定
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	svc := s3.New(sess)
	return &S3Uploader{svc: svc}, nil
}

func (uploader *S3Uploader) Upload(bucketName string, image []byte, key string) (string, error) {
	_, err := uploader.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(image),
	})
	fmt.Println(err)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return key, nil
}
