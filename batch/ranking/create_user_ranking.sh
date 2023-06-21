#!/bin/bash

# Set the date
DATE=$(date "+%Y%m%d")

# Set the database connection details
DB_HOST=$host
DB_PORT=$port
DB_NAME=$dbname
DB_USER=$username
DB_PASS=$password

echo DB_HOST
echo DB_PORT
echo DB_NAME
echo DB_USER
echo DB_PASS

# Create the overall ranking
psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -c "
INSERT INTO global_ranking (username, num_completed_exams, ranking_date)
SELECT username, COUNT(*), '$DATE'
FROM user_exams
WHERE is_completed = true
GROUP BY username
"

# Create the university-specific ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
INSERT INTO university_ranking (username, university, num_completed_exams, ranking_date)
SELECT username, university, COUNT(*), '$DATE'
FROM user_exams
WHERE is_completed = true
GROUP BY username, university
"
