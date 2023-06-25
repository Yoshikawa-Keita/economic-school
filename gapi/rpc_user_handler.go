package gapi

import (
	"context"
	"database/sql"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/service"
	"economic-school/util"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

func (server *Server) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameParam) (*pb.User, error) {

	username := req.GetUsername()
	if username == "" {
		return nil, status.Error(codes.InvalidArgument, "username cannot be empty")
	}

	user, err := server.store.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "an unexpected error occurred: %v", err)
	}

	return convertUser(user), nil
}

func (server *Server) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailParam) (*pb.User, error) {

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

	return convertUser(user), nil
}

func (server *Server) SendPasswordResetEmail(ctx context.Context, req *pb.SendPasswordResetEmailRequest) error {

	email := req.GetEmail()
	if email == "" {
		return status.Error(codes.InvalidArgument, "email cannot be empty")
	}

	user, err := server.store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Error(codes.NotFound, "user not found")
		}
		return status.Errorf(codes.Internal, "an unexpected error occurred: %v", err)
	}

	passwordResetEmail, err := server.store.CreatePasswordResetEmail(ctx, db.CreatePasswordResetEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("failed to create password reset email: %w", err)
	}

	message := map[string]string{
		"Username":   user.Username,
		"Email":      email,
		"Email_id":   strconv.FormatInt(passwordResetEmail.ID, 10),
		"SecretCode": passwordResetEmail.SecretCode,
	}

	config, err := util.LoadConfig("..")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = service.SendMessageToSQS(config, config.SQSEmailSendingQueue, message)
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %w", err)
	}
	log.Info().Str("email", email).Msg("sending password reset message to email-sending-queue")

	return nil
}
