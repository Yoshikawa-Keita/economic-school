package gapi

import (
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		UserType:          user.UserType,
		ProfileImageUrl:   user.ProfileImageUrl,
		IsEmailVerified:   user.IsEmailVerified,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		Version:           user.Version,
	}
}

func convertExam(dbExam db.Exam) *pb.Exam {
	return &pb.Exam{
		ExamId:         dbExam.ExamID,
		University:     dbExam.University,
		Subject:        dbExam.Subject,
		Year:           dbExam.Year,
		QuestionNum:    dbExam.QuestionNum,
		QuestionPdfUrl: dbExam.QuestionPdfUrl.String,
		AnswerPdfUrl:   dbExam.AnswerPdfUrl.String,
		VideoUrl:       dbExam.VideoUrl.String,
		CritiqueUrl:    dbExam.CritiqueUrl.String,
	}
}

func convertUserExam(userExam db.UserExam) *pb.UserExam {

	return &pb.UserExam{
		Username:    userExam.Username,
		ExamId:      userExam.ExamID,
		University:  userExam.University,
		IsCompleted: userExam.IsCompleted,
		CompletedAt: timestamppb.New(userExam.CompletedAt.Time),
	}
}
