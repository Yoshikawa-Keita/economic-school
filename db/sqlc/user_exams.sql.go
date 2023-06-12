// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user_exams.sql

package db

import (
	"context"
	"database/sql"
)

const getUserExam = `-- name: GetUserExam :one
SELECT username, exam_id, university, is_completed, completed_at FROM user_exams WHERE username = $1 AND exam_id = $2 LIMIT 1
`

type GetUserExamParams struct {
	Username string `json:"username"`
	ExamID   int32  `json:"exam_id"`
}

func (q *Queries) GetUserExam(ctx context.Context, arg GetUserExamParams) (UserExam, error) {
	row := q.db.QueryRowContext(ctx, getUserExam, arg.Username, arg.ExamID)
	var i UserExam
	err := row.Scan(
		&i.Username,
		&i.ExamID,
		&i.University,
		&i.IsCompleted,
		&i.CompletedAt,
	)
	return i, err
}

const listCompletedUserExams = `-- name: ListCompletedUserExams :many
SELECT username, exam_id, university, is_completed, completed_at FROM user_exams
WHERE username = $1
AND is_completed = true
`

func (q *Queries) ListCompletedUserExams(ctx context.Context, username string) ([]UserExam, error) {
	rows, err := q.db.QueryContext(ctx, listCompletedUserExams, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserExam{}
	for rows.Next() {
		var i UserExam
		if err := rows.Scan(
			&i.Username,
			&i.ExamID,
			&i.University,
			&i.IsCompleted,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertUserExam = `-- name: UpsertUserExam :one
INSERT INTO user_exams (username, exam_id, university, is_completed, completed_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (username, exam_id) DO UPDATE
SET is_completed = EXCLUDED.is_completed, completed_at = EXCLUDED.completed_at
RETURNING username, exam_id, university, is_completed, completed_at
`

type UpsertUserExamParams struct {
	Username    string       `json:"username"`
	ExamID      int32        `json:"exam_id"`
	University  string       `json:"university"`
	IsCompleted bool         `json:"is_completed"`
	CompletedAt sql.NullTime `json:"completed_at"`
}

func (q *Queries) UpsertUserExam(ctx context.Context, arg UpsertUserExamParams) (UserExam, error) {
	row := q.db.QueryRowContext(ctx, upsertUserExam,
		arg.Username,
		arg.ExamID,
		arg.University,
		arg.IsCompleted,
		arg.CompletedAt,
	)
	var i UserExam
	err := row.Scan(
		&i.Username,
		&i.ExamID,
		&i.University,
		&i.IsCompleted,
		&i.CompletedAt,
	)
	return i, err
}
