#!/bin/bash

# Set the date
DATE_TO=$(date "+%Y%m%d")
DATE_FROM=$(date -d "$DATE_TO -7 days" "+%Y%m%d")

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
INSERT INTO global_ranking (username, num_completed_exams, ranking_date)
SELECT username, COUNT(*), '$DATE_TO'
FROM user_exams
WHERE is_completed = true
GROUP BY username
ORDER BY COUNT(*) DESC
"

# Truncate the existing university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
TRUNCATE university_ranking
"

# Create the university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
INSERT INTO university_ranking (username, university, num_completed_exams, ranking_date)
SELECT username, university, COUNT(*), '$DATE_TO'
FROM user_exams
WHERE is_completed = true
GROUP BY username, university
ORDER BY COUNT(*) DESC
"

# Truncate the existing weekly overall ranking
psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -v ON_ERROR_STOP=1 -c "
TRUNCATE weekly_global_ranking
"

# Create the weekly overall ranking
psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -v ON_ERROR_STOP=1 -c "
INSERT INTO weekly_global_ranking (username, completed_exams_count, ranking_date)
SELECT username, COUNT(*), '$DATE_TO'
FROM user_exams
WHERE is_completed = true AND date(completed_at) BETWEEN '$DATE_FROM' AND '$DATE_TO'
GROUP BY username
ORDER BY COUNT(*) DESC
"

# Truncate the existing weekly university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
TRUNCATE weekly_university_ranking
"

# Create the weekly university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -c "
INSERT INTO weekly_university_ranking (username, university, completed_exams_count, ranking_date)
SELECT username, university, COUNT(*), '$DATE_TO'
FROM user_exams
WHERE is_completed = true AND date(completed_at) BETWEEN '$DATE_FROM' AND '$DATE_TO'
GROUP BY username, university
ORDER BY COUNT(*) DESC
"
