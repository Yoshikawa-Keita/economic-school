#!/bin/bash

# Set the date
DATE_TO=$(date "+%Y%m%d")
DATE_FROM=$(date -D "%Y%m%d" -d "$DATE_TO - 7 days" "+%Y%m%d")

# Set the database connection details
DB_HOST=$host
DB_PORT=$port
DB_NAME=$dbname
DB_USER=$username
DB_PASS=$password
export PGPASSWORD=$DB_PASS

# Truncate the existing global ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
TRUNCATE global_ranking
"

# Create the overall ranking
psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -v ON_ERROR_STOP=1 -c "
INSERT INTO global_ranking (username, num_completed_exams, ranking, ranking_date)
SELECT username, COUNT(*), RANK() OVER (ORDER BY COUNT(*) DESC), '$DATE_TO'
FROM user_exams
WHERE is_completed = true
GROUP BY username
"

# Truncate the existing university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
TRUNCATE university_ranking
"

# Create the university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
INSERT INTO university_ranking (username, university, num_completed_exams, ranking, ranking_date)
SELECT username, university, COUNT(*), RANK() OVER (PARTITION BY university ORDER BY COUNT(*) DESC), '$DATE_TO'
FROM user_exams
WHERE is_completed = true
GROUP BY username, university
"

# Truncate the existing weekly overall ranking
psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -v ON_ERROR_STOP=1 -c "
TRUNCATE weekly_global_ranking
"

# Create the weekly overall ranking
psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -v ON_ERROR_STOP=1 -c "
INSERT INTO weekly_global_ranking (username, completed_exams_count, ranking, ranking_date)
SELECT username, COUNT(*), RANK() OVER (ORDER BY COUNT(*) DESC), '$DATE_TO'
FROM user_exams
WHERE is_completed = true AND completed_at BETWEEN TO_TIMESTAMP($DATE_FROM, 'YYYYMMDD') AND (TO_TIMESTAMP($DATE_TO, 'YYYYMMDD') + INTERVAL '1 day')
GROUP BY username
"

# Truncate the existing weekly university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
TRUNCATE weekly_university_ranking
"

# Create the weekly university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
INSERT INTO weekly_university_ranking (username, university, completed_exams_count, ranking, ranking_date)
SELECT username, university, COUNT(*), RANK() OVER (PARTITION BY university ORDER BY COUNT(*) DESC), '$DATE_TO'
FROM user_exams
WHERE is_completed = true AND completed_at BETWEEN TO_TIMESTAMP($DATE_FROM, 'YYYYMMDD') AND (TO_TIMESTAMP($DATE_TO, 'YYYYMMDD') + INTERVAL '1 day')
GROUP BY username, university
"
