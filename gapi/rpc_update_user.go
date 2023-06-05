package gapi

import (
	"context"
	"database/sql"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/util"
	"economic-school/val"
	"encoding/base64"
	"fmt"
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
