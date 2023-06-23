-- name: GetGlobalRanking :many
SELECT * FROM global_ranking
ORDER BY num_completed_exams DESC;

-- name: GetWeeklyGlobalRanking :many
SELECT * FROM weekly_global_ranking
ORDER BY completed_exams_count DESC;

-- name: GetUniversityRanking :many
SELECT * FROM university_ranking
ORDER BY university, num_completed_exams DESC;

-- name: GetWeeklyUniversityRanking :many
SELECT * FROM weekly_university_ranking
ORDER BY university, completed_exams_count DESC;

