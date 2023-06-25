-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-06-24T21:05:40.340Z

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "user_type" int NOT NULL DEFAULT 1,
  "profile_image_url" varchar NOT NULL DEFAULT 'default_profile_image.jpg',
  "is_email_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "version" int NOT NULL
);

CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '24 hours')
);

CREATE TABLE "password_reset_emails" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '24 hours')
);

CREATE TABLE "email_logs" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "bounce_count" int,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "exams" (
  "exam_id" serial PRIMARY KEY,
  "university" varchar NOT NULL,
  "subject" varchar NOT NULL,
  "year" int NOT NULL,
  "question_num" int NOT NULL,
  "question_pdf_url" varchar,
  "answer_pdf_url" varchar,
  "video_url" varchar,
  "critique_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_exams" (
  "username" varchar,
  "exam_id" int,
  "university" varchar NOT NULL,
  "is_completed" bool NOT NULL DEFAULT true,
  "completed_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("username", "exam_id")
);

CREATE TABLE "global_ranking" (
  "username" varchar PRIMARY KEY,
  "num_completed_exams" int NOT NULL,
  "ranking" int NOT NULL,
  "ranking_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "university_ranking" (
  "username" varchar,
  "university" varchar NOT NULL,
  "num_completed_exams" int NOT NULL,
  "ranking" int NOT NULL,
  "ranking_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("university", "username", "ranking_date")
);

CREATE TABLE "weekly_global_ranking" (
  "username" varchar,
  "completed_exams_count" int NOT NULL,
  "ranking" int NOT NULL,
  "ranking_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "ranking_date")
);

CREATE TABLE "weekly_university_ranking" (
  "username" varchar,
  "university" varchar NOT NULL,
  "completed_exams_count" int NOT NULL,
  "ranking" int NOT NULL,
  "ranking_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("university", "username", "ranking_date")
);

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "password_reset_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_exams" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_exams" ADD FOREIGN KEY ("exam_id") REFERENCES "exams" ("exam_id");

ALTER TABLE "global_ranking" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "university_ranking" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "weekly_global_ranking" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "weekly_university_ranking" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
