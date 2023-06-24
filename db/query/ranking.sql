-- name: GetGlobalRanking :many
SELECT * FROM global_ranking
ORDER BY ranking;

-- name: GetWeeklyGlobalRanking :many
SELECT * FROM weekly_global_ranking
ORDER BY ranking;

-- name: GetUniversityRanking :many
SELECT * FROM university_ranking
ORDER BY university, ranking;

-- name: GetWeeklyUniversityRanking :many
SELECT * FROM weekly_university_ranking
ORDER BY university, ranking;
