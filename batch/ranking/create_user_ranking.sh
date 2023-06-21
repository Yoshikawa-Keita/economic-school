#!/bin/bash

# Set the date
DATE=$(date "+%Y%m%d")

# Set the database connection details
DB_HOST=$host # update here
DB_PORT="your_db_port" # update if you have a secret for this
DB_NAME=$dbname # update here
DB_USER=$username # update here
DB_PASS=$password # update here

# Create the overall ranking
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
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
