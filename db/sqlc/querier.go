// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateExam(ctx context.Context, arg CreateExamParams) (Exam, error)
	CreatePasswordResetEmail(ctx context.Context, arg CreatePasswordResetEmailParams) (PasswordResetEmail, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteExam(ctx context.Context, examID int32) error
	DeleteUser(ctx context.Context, username string) (User, error)
	GetExam(ctx context.Context, examID int32) (Exam, error)
	GetExamCountByUniversity(ctx context.Context) ([]GetExamCountByUniversityRow, error)
	GetGlobalRanking(ctx context.Context) ([]GlobalRanking, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUniversityRanking(ctx context.Context) ([]UniversityRanking, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserExam(ctx context.Context, arg GetUserExamParams) (UserExam, error)
	GetWeeklyGlobalRanking(ctx context.Context) ([]WeeklyGlobalRanking, error)
	GetWeeklyUniversityRanking(ctx context.Context) ([]WeeklyUniversityRanking, error)
	ListCompletedUserExams(ctx context.Context, username string) ([]UserExam, error)
	ListExams(ctx context.Context, arg ListExamsParams) ([]Exam, error)
	ListUniversities(ctx context.Context) ([]string, error)
	// -- name: ListExams :many
	// SELECT *
	// FROM exams
	// WHERE
	//     university = sqlc.narg('university')
	//     AND subject = sqlc.narg('subject')
	//     AND year = sqlc.narg('year')
	// ORDER BY exam_id;
	UpdateExam(ctx context.Context, arg UpdateExamParams) (Exam, error)
	UpdatePasswordResetEmail(ctx context.Context, arg UpdatePasswordResetEmailParams) (PasswordResetEmail, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
	UpsertUserExam(ctx context.Context, arg UpsertUserExamParams) (UserExam, error)
}

var _ Querier = (*Queries)(nil)
