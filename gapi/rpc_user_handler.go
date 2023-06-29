package gapi

import (
	"context"
	"database/sql"
	"economic-school/pb"
	"economic-school/service"
	"economic-school/util"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (server *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	username := req.GetUsername()
	if username == "" {
		return nil, status.Error(codes.InvalidArgument, "username cannot be empty")
	}

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot delete other user's info")
	}

	user, err := server.store.DeleteUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "an unexpected error occurred: %v", err)
	}

	message := map[string]string{
		"EmailType": "DELETE_USER",
		"Username":  user.Username,
		"Email":     user.Email,
	}

	config, err := util.LoadConfig("..")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load config: %v", err)
	}

	err = service.SendMessageToSQS(config, config.SQSEmailSendingQueue, message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send message to SQS: %v", err)
	}
	log.Info().Str("email", user.Email).Msg("sending password reset message to email-sending-queue")

	return &empty.Empty{}, nil
}
