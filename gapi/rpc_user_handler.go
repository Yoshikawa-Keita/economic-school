package gapi

import (
	"context"
	"database/sql"
	"economic-school/pb"
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
