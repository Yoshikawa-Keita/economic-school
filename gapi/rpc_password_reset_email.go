package gapi

import (
	"context"
	"database/sql"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/service"
	"economic-school/util"
	"economic-school/val"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (server *Server) SendPasswordResetEmail(ctx context.Context, req *pb.SendPasswordResetEmailRequest) (*emptypb.Empty, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, status.Error(codes.InvalidArgument, "email cannot be empty")
	}

	user, err := server.store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "an unexpected error occurred: %v", err)
	}

	passwordResetEmail, err := server.store.CreatePasswordResetEmail(ctx, db.CreatePasswordResetEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(32),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create password reset email: %v", err)
	}

	message := map[string]string{
		"EmailType":  "CHANGE_PASSWORD",
		"Username":   user.Username,
		"Email":      email,
		"Email_id":   strconv.FormatInt(passwordResetEmail.ID, 10),
		"SecretCode": passwordResetEmail.SecretCode,
	}

	config, err := util.LoadConfig("..")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load config: %v", err)
	}

	err = service.SendMessageToSQS(config, config.SQSEmailSendingQueue, message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send message to SQS: %v", err)
	}
	log.Info().Str("email", email).Msg("sending password reset message to email-sending-queue")

	return &emptypb.Empty{}, nil
}

func (server *Server) PasswordResetEmail(ctx context.Context, req *pb.PasswordResetEmailRequest) (*pb.PasswordResetEmailResponse, error) {
	violations := validatePasswordResetEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	txResult, err := server.store.PasswordResetEmailTx(ctx, db.PasswordResetEmailTxParams{
		EmailId:        req.GetEmailId(),
		SecretCode:     req.GetSecretCode(),
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to reset password")
	}

	rsp := &pb.PasswordResetEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	return rsp, nil
}

func validatePasswordResetEmailRequest(req *pb.PasswordResetEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, fieldViolation("email_id", err))
	}

	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err))
	}

	return violations
}
