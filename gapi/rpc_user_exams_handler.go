package gapi

import (
	"context"
	"database/sql"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) UpsertUserExam(ctx context.Context, req *pb.UpsertUserExamRequest) (*pb.UserExam, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	arg := db.UpsertUserExamParams{
		Username:    req.GetUsername(),
		ExamID:      req.GetExamId(),
		University:  req.GetUniversity(),
		IsCompleted: req.GetIsCompleted(),
		CompletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: req.GetIsCompleted(),
		},
	}

	userExam, err := server.store.UpsertUserExam(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user exam: %s", err)
	}
	rsp := convertUserExam(userExam)
	return rsp, nil
}

func (server *Server) GetUserExam(ctx context.Context, req *pb.GetUserExamRequest) (*pb.UserExam, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	arg := db.GetUserExamParams{
		Username: req.GetUsername(),
		ExamID:   req.GetExamId(),
	}
	userExam, err := server.store.GetUserExam(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user exam not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get user exam: %s", err)
	}
	rsp := convertUserExam(userExam)
	return rsp, err
}

func (server *Server) ListCompletedUserExams(ctx context.Context, req *pb.ListCompletedUserExamsRequest) (*pb.ListCompletedUserExamsResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	//arg := db.ListCompletedUserExamsParams{
	//	Username:   req.GetUsername()
	//	//University: req.GetUniversity(),
	//}

	userExams, err := server.store.ListCompletedUserExams(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list completed user exams: %s", err)
	}
	responseExams := make([]*pb.UserExam, len(userExams))
	for i, userExam := range userExams {
		responseExams[i] = convertUserExam(userExam)
	}

	return &pb.ListCompletedUserExamsResponse{UserExams: responseExams}, nil
}
