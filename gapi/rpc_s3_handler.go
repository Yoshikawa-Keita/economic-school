package gapi

import (
	"context"
	"economic-school/pb"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) GetSignedUrl(ctx context.Context, req *pb.GetSignedUrlRequest) (*pb.GetSignedUrlResponse, error) {
	// ユーザーの認証
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// リクエストの検証
	violations := validateGetSignedURLRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// ファイル名を設定
	filename := req.GetFilePath()

	//// 署名付きURLを生成
	//config, err := util.LoadConfig("..")
	//signedURL, err := server.s3Uploader.GetSignedURL(config.S3BucketNameTransferExam, filename, 1*time.Hour)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to generate signed URL: %s", err)
	//}

	// 署名付きURLを生成
	//config, err := util.LoadConfig("..")
	signedURL, err := server.cloudFrontSigner.GetSignedURL(filename, 1*time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate signed URL: %s", err)
	}
	rsp := &pb.GetSignedUrlResponse{
		SignedUrl: signedURL,
	}

	return rsp, nil
}

func validateGetSignedURLRequest(req *pb.GetSignedUrlRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// ファイル名の検証
	if req.GetFilePath() == "" {
		violations = append(violations, fieldViolation("filePath", fmt.Errorf("file empty")))
	}

	return violations
}
