-- name: CreateExam :one
INSERT INTO exams (
    university, subject, year, question_num, question_pdf_url, answer_pdf_url, video_url, critique_url
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         ) RETURNING *;

-- name: GetExam :one
SELECT * FROM exams WHERE exam_id = $1;

-- name: ListExams :many
SELECT *
FROM exams
WHERE
    (university = coalesce(sqlc.narg('university'), university))
  AND (subject = coalesce(sqlc.narg('subject'), subject))
  AND (year = coalesce(sqlc.narg('year'), year))
ORDER BY year DESC, question_num ASC;


-- -- name: ListExams :many
-- SELECT *
-- FROM exams
-- WHERE
--     university = sqlc.narg('university')
--     AND subject = sqlc.narg('subject')
--     AND year = sqlc.narg('year')
-- ORDER BY exam_id;


-- name: UpdateExam :one
UPDATE exams SET
                 university = COALESCE($2, university),
                 subject = COALESCE($3, subject),
                 year = COALESCE($4, year),
                 question_num = COALESCE($5, question_num),
                 question_pdf_url = COALESCE($6, question_pdf_url),
                 answer_pdf_url = COALESCE($7, answer_pdf_url),
                 video_url = COALESCE($8, video_url),
                 critique_url = COALESCE($9, critique_url)
WHERE exam_id = $1
    RETURNING *;
-- name: DeleteExam :exec
DELETE FROM exams WHERE exam_id = $1;
