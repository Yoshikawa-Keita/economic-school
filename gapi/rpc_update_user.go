package gapi

import (
	"context"
	"database/sql"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/service"
	"economic-school/util"
	"economic-school/val"
	"encoding/base64"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	arg := db.UpdateUserTxParams{
		UpdateUserParams: db.UpdateUserParams{
			Username: req.GetUsername(),
			FullName: sql.NullString{
				String: req.GetFullName(),
				Valid:  req.FullName != nil,
			},
			Email: sql.NullString{
				String: req.GetEmail(),
				Valid:  req.Email != nil,
			},
			ProfileImageUrl: sql.NullString{
				String: req.GetUsername() + ".jpg",
				Valid:  req.ProfileImage != nil,
			},
		},
		UpdateImage: func(user db.User) error {
			if req.GetProfileImage() == "" {
				return nil
			}
			profileImageBase64 := req.GetProfileImage()
			profileImage, err := base64.StdEncoding.DecodeString(profileImageBase64)
			if err != nil {
				return fmt.Errorf("failed to decode profile image: %w", err)
			}
			// Upload profile image to S3
			key := req.Username + ".jpg"
			config, err := util.LoadConfig("..")
			_, uploadErr := server.s3Uploader.Upload(config.S3BucketNameUserProfile, profileImage, key)
			if uploadErr != nil {
				return fmt.Errorf("failed to upload profile image: %w", err)
			}
			return nil
		},
	}

	if req.Password != nil {
		hashedPassword, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}

		arg.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  true,
		}

		arg.PasswordChangedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	txResult, err := server.store.UpdateUserTx(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(txResult.User),
	}
	return rsp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	if req.FullName != nil {
		if err := val.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}

	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	return violations
}

func (server *Server) UpdateUserEmail(ctx context.Context, req *pb.UpdateUserEmailRequest) (*pb.UpdateUserEmailResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateUserEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	_, err = server.store.GetUserByEmail(ctx, req.GetEmail())
	if err == nil {
		return nil, fmt.Errorf("email already registered")
	}

	message := map[string]string{
		"EmailType": "CHANGE_EMAIL",
		"Username":  req.Username,
		"Email":     req.Email,
	}

	config, err := util.LoadConfig("..")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Send the message to SQS
	err = service.SendMessageToSQS(config, config.SQSEmailSendingQueue, message)
	if err != nil {
		return nil, fmt.Errorf("failed to send message to SQS: %w", err)
	}
	log.Info().Str("email", user.Email).Msg("sending verification message to email-sending-queue")

	rsp := &pb.UpdateUserEmailResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateUpdateUserEmailRequest(req *pb.UpdateUserEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Email != "" {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	return violations
}

func (server *Server) VerifyChangedEmail(ctx context.Context, req *pb.VerifyChangedEmailRequest) (*pb.VerifyChangedEmailResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	_, err = server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	arg := db.UpdateUserEmailTxParams{
		UpdateUserParams: db.UpdateUserParams{
			Username: req.GetUsername(),
			Email: sql.NullString{
				String: req.GetEmail(),
				Valid:  true,
			},
		},
	}

	_, err = server.store.UpdateUserEmailTx(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.VerifyChangedEmailResponse{
		IsVerified: true,
	}
	return rsp, nil
}
