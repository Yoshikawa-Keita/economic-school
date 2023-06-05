package service

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"economic-school/util"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
	"time"
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
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return key, nil
}

func (uploader *S3Uploader) GetSignedURL(bucketName string, filename string, duration time.Duration) (string, error) {
	// 署名付きURLを生成するリクエストを作成
	req, _ := uploader.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})

	// リクエストから署名付きURLを生成
	urlStr, err := req.Presign(duration)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

type CloudFrontSigner struct {
	keyID      string
	privateKey *rsa.PrivateKey
}

func NewCloudFrontSigner(config util.Config) (*CloudFrontSigner, error) {
	block, _ := pem.Decode([]byte(strings.Replace(config.CloudFrontPrivateKey, "\\n", "\n", -1)))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &CloudFrontSigner{keyID: config.CloudFrontKeyId, privateKey: privateKey}, nil
}

func (signer *CloudFrontSigner) GetSignedURL(url string, duration time.Duration) (string, error) {
	urlSigner := sign.NewURLSigner(signer.keyID, signer.privateKey)

	signedURL, err := urlSigner.Sign(url, time.Now().Add(duration))
	if err != nil {
		return "", fmt.Errorf("failed to sign url: %w", err)
	}

	return signedURL, nil
}
