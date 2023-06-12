package gapi

import (
	"context"
	"database/sql"
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/val"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) CreateExam(ctx context.Context, req *pb.CreateExamRequest) (*pb.Exam, error) {
	violations := validateCreateExamRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg := db.CreateExamTxParams{
		CreateExamParams: db.CreateExamParams{
			University:  req.GetUniversity(),
			Subject:     req.GetSubject(),
			Year:        req.GetYear(),
			QuestionNum: req.GetQuestionNum(),
			QuestionPdfUrl: sql.NullString{
				String: req.GetQuestionPdfUrl(),
				Valid:  req.GetQuestionPdfUrl() != "",
			},
			AnswerPdfUrl: sql.NullString{
				String: req.GetAnswerPdfUrl(),
				Valid:  req.GetAnswerPdfUrl() != "",
			},
			VideoUrl: sql.NullString{
				String: req.GetVideoUrl(),
				Valid:  req.GetVideoUrl() != "",
			},
		},
	}

	result, err := server.store.CreateExamTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "exam already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create exam: %s", err)
	}
	rsp := convertExam(result.Exam)
	return rsp, nil
}

func validateCreateExamRequest(req *pb.CreateExamRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Here you can add field-specific validations, similar to validateCreateUserRequest
	return violations
}

func (server *Server) GetExam(ctx context.Context, req *pb.GetExamRequest) (*pb.Exam, error) {
	exam, err := server.store.GetExam(ctx, req.GetExamId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "exam not found: %s", err)
	}
	rsp := convertExam(exam)
	return rsp, nil
}

func (server *Server) ListExams(ctx context.Context, req *pb.ListExamsRequest) (*pb.ListExamsResponse, error) {
	arg := db.ListExamsParams{
		University: sql.NullString{
			String: req.GetUniversity(),
			Valid:  req.University != nil && *req.University != "",
		},
		Subject: sql.NullString{
			String: req.GetSubject(),
			Valid:  req.Subject != nil && *req.Subject != "",
		},
		Year: sql.NullInt32{
			Int32: req.GetYear(),
			Valid: req.Year != nil && *req.Year != 0,
		},
	}
	exams, err := server.store.ListExams(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list exams: %s", err)
	}

	rsp := &pb.ListExamsResponse{}
	for _, exam := range exams {
		rsp.Exams = append(rsp.Exams, convertExam(exam))
	}
	return rsp, nil
}

func (server *Server) UpdateExam(ctx context.Context, req *pb.UpdateExamRequest) (*pb.Exam, error) {
	violations := validateUpdateExamRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateExamParams{
		ExamID:      req.GetExam().GetExamId(),
		University:  req.GetExam().GetUniversity(),
		Subject:     req.GetExam().GetSubject(),
		Year:        req.GetExam().GetYear(),
		QuestionNum: req.GetExam().GetQuestionNum(),
		QuestionPdfUrl: sql.NullString{
			String: req.GetExam().GetQuestionPdfUrl(),
			Valid:  req.GetExam().GetQuestionPdfUrl() != "",
		},
		AnswerPdfUrl: sql.NullString{
			String: req.GetExam().GetAnswerPdfUrl(),
			Valid:  req.GetExam().GetAnswerPdfUrl() != "",
		},
		VideoUrl: sql.NullString{
			String: req.GetExam().GetVideoUrl(),
			Valid:  req.GetExam().GetVideoUrl() != "",
		},
	}

	exam, err := server.store.UpdateExam(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update exam: %s", err)
	}
	rsp := convertExam(exam)
	return rsp, nil
}

func (server *Server) DeleteExam(ctx context.Context, req *pb.DeleteExamRequest) (*empty.Empty, error) {
	err := server.store.DeleteExam(ctx, req.GetExamId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete exam: %s", err)
	}
	return &empty.Empty{}, nil
}

func (server *Server) ListUniversities(ctx context.Context, req *emptypb.Empty) (*pb.ListUniversitiesResponse, error) {
	universities, err := server.store.ListUniversities(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list distinct universities: %s", err)
	}

	rsp := &pb.ListUniversitiesResponse{
		Universities: universities,
	}
	return rsp, nil
}

func (server *Server) GetExamCountByUniversity(ctx context.Context, req *emptypb.Empty) (*pb.GetExamCountByUniversityResponse, error) {
	examCounts, err := server.store.GetExamCountByUniversity(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get exam count by university: %s", err)
	}

	rsp := &pb.GetExamCountByUniversityResponse{
		ExamCountByUniversity: make([]*pb.ExamCountByUniversity, len(examCounts)),
	}
	for i, count := range examCounts {
		rsp.ExamCountByUniversity[i] = &pb.ExamCountByUniversity{
			University: count.University,
			Count:      count.Count,
		}
	}
	return rsp, nil
}

func validateUpdateExamRequest(req *pb.UpdateExamRequest) []*errdetails.BadRequest_FieldViolation {
	var violations []*errdetails.BadRequest_FieldViolation

	if req.GetExam().GetUniversity() == "" {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "university",
			Description: "University name is required",
		})
	}

	if req.GetExam().GetSubject() == "" {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "subject",
			Description: "Subject name is required",
		})
	}

	if req.GetExam().GetYear() == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "year",
			Description: "Year is required",
		})
	}

	if req.GetExam().GetQuestionNum() == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "question_num",
			Description: "Question number is required",
		})
	}

	if req.GetExam().GetQuestionPdfUrl() != "" {
		if err := val.ValidateUrl(req.GetExam().GetQuestionPdfUrl()); err != nil {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field:       "question_pdf_url",
				Description: "Invalid URL format",
			})
		}
	}

	if req.GetExam().GetAnswerPdfUrl() != "" {
		if err := val.ValidateUrl(req.GetExam().GetAnswerPdfUrl()); err != nil {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field:       "answer_pdf_url",
				Description: "Invalid URL format",
			})
		}
	}

	if req.GetExam().GetVideoUrl() != "" {
		if err := val.ValidateUrl(req.GetExam().GetVideoUrl()); err != nil {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field:       "video_url",
				Description: "Invalid URL format",
			})
		}
	}

	return violations
}
