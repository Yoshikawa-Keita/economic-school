package gapi

import (
	"bytes"
	"context"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/util"
	"economic-school/val"
	"economic-school/worker"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	profileImageUrl := req.GetUsername() + ".jpg"
	if req.GetProfileFile() == "" {
		profileImageUrl = "default_image.png"
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:        req.GetUsername(),
			HashedPassword:  hashedPassword,
			FullName:        req.GetFullName(),
			UserType:        1,
			Email:           req.GetEmail(),
			ProfileImageUrl: profileImageUrl,
		},
		UploadImage: func() error {
			if req.GetProfileFile() == "" {
				return nil
			}
			profileImageBase64 := req.GetProfileFile()
			profileImage, err := base64.StdEncoding.DecodeString(profileImageBase64)
			if err != nil {
				return fmt.Errorf("failed to decode profile image: %w", err)
			}
			// TODO 想定外の形式のファイルがきた時の対策をすべき
			//// Detect the content type of the image
			//contentType := http.DetectContentType(profileImage)
			//
			//// Get the file extension for this content type
			//extensions, err := mime.ExtensionsByType(contentType)
			//if err != nil || len(extensions) == 0 {
			//	return fmt.Errorf("could not determine file extension for content type %s: %w", contentType, err)
			//}

			// Create the key using the username and the appropriate file extension
			//key := req.Username + extensions[0]
			key := req.Username + ".jpg"

			// Upload profile image to S3
			config, err := util.LoadConfig("..")
			_, uploadErr := server.s3Uploader.Upload(config.S3BucketNameUserProfile, profileImage, key)
			if uploadErr != nil {
				return fmt.Errorf("failed to upload profile image: %w", err)
			}
			return nil
		},
		AfterCreate: func(user db.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Username,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}

			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}

	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}

func (server *Server) uploadImageToS3(region string, bucketName string, image []byte, filename string) (string, error) {
	// AWSセッションの作成
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	// S3サービスクライアントの作成
	svc := s3.New(sess)

	// バケット名とオブジェクトキーを指定してアップロード
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(image),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return filename, nil
}
