-- name: UpsertUserExam :one
INSERT INTO user_exams (username, exam_id, university, is_completed, completed_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (username, exam_id) DO UPDATE
SET is_completed = EXCLUDED.is_completed, completed_at = EXCLUDED.completed_at
RETURNING *;

-- name: GetUserExam :one
SELECT * FROM user_exams WHERE username = $1 AND exam_id = $2 LIMIT 1;

-- name: ListCompletedUserExams :many
SELECT * FROM user_exams
WHERE username = $1
AND is_completed = true;

