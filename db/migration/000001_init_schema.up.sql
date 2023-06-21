CREATE TABLE "users" (
                         "username" varchar PRIMARY KEY,
                         "hashed_password" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "user_type" int NOT NULL,
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


ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "email_logs" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_exams" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_exams" ADD FOREIGN KEY ("exam_id") REFERENCES "exams" ("exam_id");

--------- 以下で初期データをinsertする

-- user

INSERT INTO "users" (
    "username",
    "hashed_password",
    "full_name",
    "email",
    "user_type",
    "profile_image_url",
    "is_email_verified",
    "password_changed_at",
    "created_at",
    "version"
) VALUES (
             'aran',
             '$2a$10$GRAXD/JFrWwERy/s0J81ke4xod9f95jrIgQ/NeXyNKt894Cd/r/re',
             'aaaa',
             'keita_yoshi@icloud.com',
             1,
             'aran.jpg',
             true,
             '0001-01-01 00:00:00+00',
             '2023-05-28 00:05:20.738982+00',
             1
         );


-- exam

INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2017, 2, 'kobeeco_2017_2_ma_q.pdf', 'kobeeco_2017_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2021, 1, 'kobeeco_2021_1_ma_q.pdf', 'kobeeco_2021_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2012, 3, 'kobeeco_2012_3_ma_q.pdf', 'kobeeco_2012_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2022, 3, 'kobeeco_2022_3_ma_q.pdf', 'kobeeco_2022_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2015, 2, 'kobeeco_2015_2_ma_q.pdf', 'kobeeco_2015_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2014, 1, 'kobeeco_2014_1_ma_q.pdf', 'kobeeco_2014_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2018, 2, 'kobeeco_2018_2_ma_q.pdf', 'kobeeco_2018_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2023, 4, 'kobeeco_2023_4_ma_q.pdf', 'kobeeco_2023_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2013, 3, 'kobeeco_2013_3_ma_q.pdf', 'kobeeco_2013_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2013, 2, 'kobeeco_2013_2_ma_q.pdf', 'kobeeco_2013_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2016, 3, 'kobeeco_2016_3_ma_q.pdf', 'kobeeco_2016_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2021, 2, 'kobeeco_2021_2_ma_q.pdf', 'kobeeco_2021_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2021, 3, 'kobeeco_2021_3_ma_q.pdf', 'kobeeco_2021_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2020, 1, 'kobeeco_2020_1_ma_q.pdf', 'kobeeco_2020_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2011, 2, 'kobeeco_2011_2_ma_q.pdf', 'kobeeco_2011_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2011, 3, 'kobeeco_2011_3_ma_q.pdf', 'kobeeco_2011_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2023, 3, 'kobeeco_2023_3_ma_q.pdf', 'kobeeco_2023_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2018, 1, 'kobeeco_2018_1_ma_q.pdf', 'kobeeco_2018_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2013, 4, 'kobeeco_2013_4_ma_q.pdf', 'kobeeco_2013_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2022, 1, 'kobeeco_2022_1_ma_q.pdf', 'kobeeco_2022_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'macro', 2019, 2, 'kobeeco_2019_2_ma_q.pdf', 'kobeeco_2019_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2021, 4, 'kobeeco_2021_4_mi_q.pdf', 'kobeeco_2021_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2021, 5, 'kobeeco_2021_5_mi_q.pdf', 'kobeeco_2021_5_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2016, 4, 'kobeeco_2016_4_mi_q.pdf', 'kobeeco_2016_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2014, 3, 'kobeeco_2014_3_mi_q.pdf', 'kobeeco_2014_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2014, 2, 'kobeeco_2014_2_mi_q.pdf', 'kobeeco_2014_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2023, 2, 'kobeeco_2023_2_mi_q.pdf', 'kobeeco_2023_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2019, 3, 'kobeeco_2019_3_mi_q.pdf', 'kobeeco_2019_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2015, 1, 'kobeeco_2015_1_mi_q.pdf', 'kobeeco_2015_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2012, 1, 'kobeeco_2012_1_mi_q.pdf', 'kobeeco_2012_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2016, 2, 'kobeeco_2016_2_mi_q.pdf', 'kobeeco_2016_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2017, 1, 'kobeeco_2017_1_mi_q.pdf', 'kobeeco_2017_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2011, 4, 'kobeeco_2011_4_mi_q.pdf', 'kobeeco_2011_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2019, 1, 'kobeeco_2019_1_mi_q.pdf', 'kobeeco_2019_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2015, 3, 'kobeeco_2015_3_mi_q.pdf', 'kobeeco_2015_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2022, 2, 'kobeeco_2022_2_mi_q.pdf', 'kobeeco_2022_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2018, 3, 'kobeeco_2018_3_mi_q.pdf', 'kobeeco_2018_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2023, 1, 'kobeeco_2023_1_mi_q.pdf', 'kobeeco_2023_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2011, 1, 'kobeeco_2011_1_mi_q.pdf', 'kobeeco_2011_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2020, 2, 'kobeeco_2020_2_mi_q.pdf', 'kobeeco_2020_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2017, 3, 'kobeeco_2017_3_mi_q.pdf', 'kobeeco_2017_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2016, 1, 'kobeeco_2016_1_mi_q.pdf', 'kobeeco_2016_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2013, 1, 'kobeeco_2013_1_mi_q.pdf', 'kobeeco_2013_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobeeco', 'micro', 2012, 2, 'kobeeco_2012_2_mi_q.pdf', 'kobeeco_2012_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2012, 4, 'kobemng_2012_4_ma_q.pdf', 'kobemng_2012_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2014, 1, 'kobemng_2014_1_ma_q.pdf', 'kobemng_2014_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2017, 2, 'kobemng_2017_2_ma_q.pdf', 'kobemng_2017_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2013, 1, 'kobemng_2013_1_ma_q.pdf', 'kobemng_2013_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2015, 4, 'kobemng_2015_4_ma_q.pdf', 'kobemng_2015_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2016, 5, 'kobemng_2016_5_ma_q.pdf', 'kobemng_2016_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2016, 4, 'kobemng_2016_4_ma_q.pdf', 'kobemng_2016_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2014, 2, 'kobemng_2014_2_ma_q.pdf', 'kobemng_2014_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2012, 1, 'kobemng_2012_1_ma_q.pdf', 'kobemng_2012_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'macro', 2016, 3, 'kobemng_2016_3_ma_q.pdf', 'kobemng_2016_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2013, 3, 'kobemng_2013_3_mi_q.pdf', 'kobemng_2013_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2013, 2, 'kobemng_2013_2_mi_q.pdf', 'kobemng_2013_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2021, 2, 'kobemng_2021_2_mi_q.pdf', 'kobemng_2021_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2021, 3, 'kobemng_2021_3_mi_q.pdf', 'kobemng_2021_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2016, 2, 'kobemng_2016_2_mi_q.pdf', 'kobemng_2016_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2020, 1, 'kobemng_2020_1_mi_q.pdf', 'kobemng_2020_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2017, 1, 'kobemng_2017_1_mi_q.pdf', 'kobemng_2017_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2021, 5, 'kobemng_2021_5_mi_q.pdf', 'kobemng_2021_5_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2021, 4, 'kobemng_2021_4_mi_q.pdf', 'kobemng_2021_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2014, 3, 'kobemng_2014_3_mi_q.pdf', 'kobemng_2014_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2018, 1, 'kobemng_2018_1_mi_q.pdf', 'kobemng_2018_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2013, 4, 'kobemng_2013_4_mi_q.pdf', 'kobemng_2013_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2015, 1, 'kobemng_2015_1_mi_q.pdf', 'kobemng_2015_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2022, 1, 'kobemng_2022_1_mi_q.pdf', 'kobemng_2022_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2019, 2, 'kobemng_2019_2_mi_q.pdf', 'kobemng_2019_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2019, 3, 'kobemng_2019_3_mi_q.pdf', 'kobemng_2019_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2020, 3, 'kobemng_2020_3_mi_q.pdf', 'kobemng_2020_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2020, 2, 'kobemng_2020_2_mi_q.pdf', 'kobemng_2020_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2017, 2, 'kobemng_2017_2_mi_q.pdf', 'kobemng_2017_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2021, 1, 'kobemng_2021_1_mi_q.pdf', 'kobemng_2021_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2016, 1, 'kobemng_2016_1_mi_q.pdf', 'kobemng_2016_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2022, 5, 'kobemng_2022_5_mi_q.pdf', 'kobemng_2022_5_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2022, 4, 'kobemng_2022_4_mi_q.pdf', 'kobemng_2022_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2012, 2, 'kobemng_2012_2_mi_q.pdf', 'kobemng_2012_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2018, 4, 'kobemng_2018_4_mi_q.pdf', 'kobemng_2018_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2015, 3, 'kobemng_2015_3_mi_q.pdf', 'kobemng_2015_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2015, 2, 'kobemng_2015_2_mi_q.pdf', 'kobemng_2015_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2022, 2, 'kobemng_2022_2_mi_q.pdf', 'kobemng_2022_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2022, 3, 'kobemng_2022_3_mi_q.pdf', 'kobemng_2022_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2019, 1, 'kobemng_2019_1_mi_q.pdf', 'kobemng_2019_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2018, 3, 'kobemng_2018_3_mi_q.pdf', 'kobemng_2018_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2018, 2, 'kobemng_2018_2_mi_q.pdf', 'kobemng_2018_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kobemng', 'micro', 2020, 4, 'kobemng_2020_4_mi_q.pdf', 'kobemng_2020_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2014, 3, 'kyoto_2014_3_ma_q.pdf', 'kyoto_2014_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2013, 5, 'kyoto_2013_5_ma_q.pdf', 'kyoto_2013_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2013, 4, 'kyoto_2013_4_ma_q.pdf', 'kyoto_2013_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2016, 5, 'kyoto_2016_5_ma_q.pdf', 'kyoto_2016_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2016, 4, 'kyoto_2016_4_ma_q.pdf', 'kyoto_2016_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2021, 4, 'kyoto_2021_4_ma_q.pdf', 'kyoto_2021_4_ma_a.pdf', '', 'kyoto_2021_4_ma_qr.pdf');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2021, 5, 'kyoto_2021_5_ma_q.pdf', 'kyoto_2021_5_ma_a.pdf', '', 'kyoto_2021_5_ma_qr.pdf');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2011, 3, 'kyoto_2011_3_ma_q.pdf', 'kyoto_2011_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2017, 7, 'kyoto_2017_7_ma_q.pdf', 'kyoto_2017_7_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2017, 6, 'kyoto_2017_6_ma_q.pdf', 'kyoto_2017_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2020, 6, 'kyoto_2020_6_ma_q.pdf', 'kyoto_2020_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2021, 3, 'kyoto_2021_3_ma_q.pdf', 'kyoto_2021_3_ma_a.pdf', '', 'kyoto_2021_3_ma_qr.pdf');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2011, 4, 'kyoto_2011_4_ma_q.pdf', 'kyoto_2011_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2011, 5, 'kyoto_2011_5_ma_q.pdf', 'kyoto_2011_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2014, 4, 'kyoto_2014_4_ma_q.pdf', 'kyoto_2014_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2014, 5, 'kyoto_2014_5_ma_q.pdf', 'kyoto_2014_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2019, 4, 'kyoto_2019_4_ma_q.pdf', 'kyoto_2019_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2019, 5, 'kyoto_2019_5_ma_q.pdf', 'kyoto_2019_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2013, 3, 'kyoto_2013_3_ma_q.pdf', 'kyoto_2013_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2017, 4, 'kyoto_2017_4_ma_q.pdf', 'kyoto_2017_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2017, 5, 'kyoto_2017_5_ma_q.pdf', 'kyoto_2017_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2020, 5, 'kyoto_2020_5_ma_q.pdf', 'kyoto_2020_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2020, 4, 'kyoto_2020_4_ma_q.pdf', 'kyoto_2020_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2010, 2, 'kyoto_2010_2_ma_q.pdf', 'kyoto_2010_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2021, 6, 'kyoto_2021_6_ma_q.pdf', 'kyoto_2021_6_ma_a.pdf', '', 'kyoto_2021_6_ma_qr.pdf');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2015, 3, 'kyoto_2015_3_ma_q.pdf', 'kyoto_2015_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2012, 4, 'kyoto_2012_4_ma_q.pdf', 'kyoto_2012_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2012, 5, 'kyoto_2012_5_ma_q.pdf', 'kyoto_2012_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2015, 5, 'kyoto_2015_5_ma_q.pdf', 'kyoto_2015_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2015, 4, 'kyoto_2015_4_ma_q.pdf', 'kyoto_2015_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2019, 6, 'kyoto_2019_6_ma_q.pdf', 'kyoto_2019_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2018, 5, 'kyoto_2018_5_ma_q.pdf', 'kyoto_2018_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2018, 4, 'kyoto_2018_4_ma_q.pdf', 'kyoto_2018_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'macro', 2017, 3, 'kyoto_2017_3_ma_q.pdf', 'kyoto_2017_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2013, 1, 'kyoto_2013_1_mi_q.pdf', 'kyoto_2013_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2012, 2, 'kyoto_2012_2_mi_q.pdf', 'kyoto_2012_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2012, 3, 'kyoto_2012_3_mi_q.pdf', 'kyoto_2012_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2020, 3, 'kyoto_2020_3_mi_q.pdf', 'kyoto_2020_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2020, 2, 'kyoto_2020_2_mi_q.pdf', 'kyoto_2020_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2017, 2, 'kyoto_2017_2_mi_q.pdf', 'kyoto_2017_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2021, 1, 'kyoto_2021_1_mi_q.pdf', 'kyoto_2021_1_mi_a.pdf', '', 'kyoto_2021_1_mi_qr.pdf');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2016, 1, 'kyoto_2016_1_mi_q.pdf', 'kyoto_2016_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2011, 1, 'kyoto_2011_1_mi_q.pdf', 'kyoto_2011_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2015, 2, 'kyoto_2015_2_mi_q.pdf', 'kyoto_2015_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2019, 1, 'kyoto_2019_1_mi_q.pdf', 'kyoto_2019_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2014, 1, 'kyoto_2014_1_mi_q.pdf', 'kyoto_2014_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2018, 3, 'kyoto_2018_3_mi_q.pdf', 'kyoto_2018_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2018, 2, 'kyoto_2018_2_mi_q.pdf', 'kyoto_2018_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2021, 2, 'kyoto_2021_2_mi_q.pdf', 'kyoto_2021_2_mi_a.pdf', '', 'kyoto_2021_2_mi_qr.pdf');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2016, 3, 'kyoto_2016_3_mi_q.pdf', 'kyoto_2016_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2016, 2, 'kyoto_2016_2_mi_q.pdf', 'kyoto_2016_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2020, 1, 'kyoto_2020_1_mi_q.pdf', 'kyoto_2020_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2017, 1, 'kyoto_2017_1_mi_q.pdf', 'kyoto_2017_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2012, 1, 'kyoto_2012_1_mi_q.pdf', 'kyoto_2012_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2013, 2, 'kyoto_2013_2_mi_q.pdf', 'kyoto_2013_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2014, 2, 'kyoto_2014_2_mi_q.pdf', 'kyoto_2014_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2018, 1, 'kyoto_2018_1_mi_q.pdf', 'kyoto_2018_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2015, 1, 'kyoto_2015_1_mi_q.pdf', 'kyoto_2015_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2019, 2, 'kyoto_2019_2_mi_q.pdf', 'kyoto_2019_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2019, 3, 'kyoto_2019_3_mi_q.pdf', 'kyoto_2019_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2010, 1, 'kyoto_2010_1_mi_q.pdf', 'kyoto_2010_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('kyoto', 'micro', 2011, 2, 'kyoto_2011_2_mi_q.pdf', 'kyoto_2011_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'macro', 2018, 1, 'nagoya_2018_1_ma_q.pdf', 'nagoya_2018_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'macro', 2015, 1, 'nagoya_2015_1_ma_q.pdf', 'nagoya_2015_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'macro', 2013, 1, 'nagoya_2013_1_ma_q.pdf', 'nagoya_2013_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'micro', 2020, 2, 'nagoya_2020_2_mi_q.pdf', 'nagoya_2020_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'micro', 2016, 1, 'nagoya_2016_1_mi_q.pdf', 'nagoya_2016_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'micro', 2019, 1, 'nagoya_2019_1_mi_q.pdf', 'nagoya_2019_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'micro', 2014, 1, 'nagoya_2014_1_mi_q.pdf', 'nagoya_2014_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'micro', 2017, 1, 'nagoya_2017_1_mi_q.pdf', 'nagoya_2017_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('nagoya', 'micro', 2020, 1, 'nagoya_2020_1_mi_q.pdf', 'nagoya_2020_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2021, 1, 'niigata_2021_1_ma_q.pdf', 'niigata_2021_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2016, 1, 'niigata_2016_1_ma_q.pdf', 'niigata_2016_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2020, 2, 'niigata_2020_2_ma_q.pdf', 'niigata_2020_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2014, 1, 'niigata_2014_1_ma_q.pdf', 'niigata_2014_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2015, 2, 'niigata_2015_2_ma_q.pdf', 'niigata_2015_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2020, 1, 'niigata_2020_1_ma_q.pdf', 'niigata_2020_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2016, 2, 'niigata_2016_2_ma_q.pdf', 'niigata_2016_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2016, 3, 'niigata_2016_3_ma_q.pdf', 'niigata_2016_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2014, 4, 'niigata_2014_4_ma_q.pdf', 'niigata_2014_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2015, 1, 'niigata_2015_1_ma_q.pdf', 'niigata_2015_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2014, 3, 'niigata_2014_3_ma_q.pdf', 'niigata_2014_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'macro', 2014, 2, 'niigata_2014_2_ma_q.pdf', 'niigata_2014_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2016, 4, 'niigata_2016_4_mi_q.pdf', 'niigata_2016_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2016, 5, 'niigata_2016_5_mi_q.pdf', 'niigata_2016_5_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2014, 5, 'niigata_2014_5_mi_q.pdf', 'niigata_2014_5_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2016, 6, 'niigata_2016_6_mi_q.pdf', 'niigata_2016_6_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2015, 3, 'niigata_2015_3_mi_q.pdf', 'niigata_2015_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2014, 6, 'niigata_2014_6_mi_q.pdf', 'niigata_2014_6_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2015, 4, 'niigata_2015_4_mi_q.pdf', 'niigata_2015_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2021, 1, 'niigata_2021_1_mi_q.pdf', 'niigata_2021_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('niigata', 'micro', 2020, 3, 'niigata_2020_3_mi_q.pdf', 'niigata_2020_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2014, 4, 'osaka_2014_4_ma_q.pdf', 'osaka_2014_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2021, 3, 'osaka_2021_3_ma_q.pdf', 'osaka_2021_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2020, 6, 'osaka_2020_6_ma_q.pdf', 'osaka_2020_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2020, 7, 'osaka_2020_7_ma_q.pdf', 'osaka_2020_7_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2021, 4, 'osaka_2021_4_ma_q.pdf', 'osaka_2021_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2021, 5, 'osaka_2021_5_ma_q.pdf', 'osaka_2021_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2016, 5, 'osaka_2016_5_ma_q.pdf', 'osaka_2016_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2016, 4, 'osaka_2016_4_ma_q.pdf', 'osaka_2016_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2013, 5, 'osaka_2013_5_ma_q.pdf', 'osaka_2013_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2013, 4, 'osaka_2013_4_ma_q.pdf', 'osaka_2013_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2012, 6, 'osaka_2012_6_ma_q.pdf', 'osaka_2012_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2014, 3, 'osaka_2014_3_ma_q.pdf', 'osaka_2014_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2014, 2, 'osaka_2014_2_ma_q.pdf', 'osaka_2014_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2017, 3, 'osaka_2017_3_ma_q.pdf', 'osaka_2017_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2015, 5, 'osaka_2015_5_ma_q.pdf', 'osaka_2015_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2015, 4, 'osaka_2015_4_ma_q.pdf', 'osaka_2015_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2012, 4, 'osaka_2012_4_ma_q.pdf', 'osaka_2012_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2012, 5, 'osaka_2012_5_ma_q.pdf', 'osaka_2012_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2016, 6, 'osaka_2016_6_ma_q.pdf', 'osaka_2016_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2020, 5, 'osaka_2020_5_ma_q.pdf', 'osaka_2020_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2020, 4, 'osaka_2020_4_ma_q.pdf', 'osaka_2020_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2017, 4, 'osaka_2017_4_ma_q.pdf', 'osaka_2017_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('osaka', 'macro', 2017, 5, 'osaka_2017_5_ma_q.pdf', 'osaka_2017_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2015, 1, 'tohoku_2015_1_ma_q.pdf', 'tohoku_2015_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2021, 2, 'tohoku_2021_2_ma_q.pdf', 'tohoku_2021_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2021, 3, 'tohoku_2021_3_ma_q.pdf', 'tohoku_2021_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2016, 3, 'tohoku_2016_3_ma_q.pdf', 'tohoku_2016_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2016, 2, 'tohoku_2016_2_ma_q.pdf', 'tohoku_2016_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2019, 4, 'tohoku_2019_4_ma_q.pdf', 'tohoku_2019_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2020, 4, 'tohoku_2020_4_ma_q.pdf', 'tohoku_2020_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2018, 2, 'tohoku_2018_2_ma_q.pdf', 'tohoku_2018_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2013, 1, 'tohoku_2013_1_ma_q.pdf', 'tohoku_2013_1_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2020, 3, 'tohoku_2020_3_ma_q.pdf', 'tohoku_2020_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'macro', 2017, 2, 'tohoku_2017_2_ma_q.pdf', 'tohoku_2017_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2012, 2, 'tohoku_2012_2_mi_q.pdf', 'tohoku_2012_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2016, 1, 'tohoku_2016_1_mi_q.pdf', 'tohoku_2016_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2021, 1, 'tohoku_2021_1_mi_q.pdf', 'tohoku_2021_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2020, 2, 'tohoku_2020_2_mi_q.pdf', 'tohoku_2020_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2014, 1, 'tohoku_2014_1_mi_q.pdf', 'tohoku_2014_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2019, 1, 'tohoku_2019_1_mi_q.pdf', 'tohoku_2019_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2015, 2, 'tohoku_2015_2_mi_q.pdf', 'tohoku_2015_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2015, 3, 'tohoku_2015_3_mi_q.pdf', 'tohoku_2015_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2017, 1, 'tohoku_2017_1_mi_q.pdf', 'tohoku_2017_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2020, 1, 'tohoku_2020_1_mi_q.pdf', 'tohoku_2020_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2012, 1, 'tohoku_2012_1_mi_q.pdf', 'tohoku_2012_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2019, 3, 'tohoku_2019_3_mi_q.pdf', 'tohoku_2019_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2019, 2, 'tohoku_2019_2_mi_q.pdf', 'tohoku_2019_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2018, 1, 'tohoku_2018_1_mi_q.pdf', 'tohoku_2018_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2014, 2, 'tohoku_2014_2_mi_q.pdf', 'tohoku_2014_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('tohoku', 'micro', 2021, 4, 'tohoku_2021_4_mi_q.pdf', 'tohoku_2021_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2017, 5, 'yokohama_2017_5_ma_q.pdf', 'yokohama_2017_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2017, 4, 'yokohama_2017_4_ma_q.pdf', 'yokohama_2017_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2022, 3, 'yokohama_2022_3_ma_q.pdf', 'yokohama_2022_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2015, 3, 'yokohama_2015_3_ma_q.pdf', 'yokohama_2015_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2018, 3, 'yokohama_2018_3_ma_q.pdf', 'yokohama_2018_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2012, 4, 'yokohama_2012_4_ma_q.pdf', 'yokohama_2012_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2015, 4, 'yokohama_2015_4_ma_q.pdf', 'yokohama_2015_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2014, 6, 'yokohama_2014_6_ma_q.pdf', 'yokohama_2014_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2012, 3, 'yokohama_2012_3_ma_q.pdf', 'yokohama_2012_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2020, 2, 'yokohama_2020_2_ma_q.pdf', 'yokohama_2020_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2010, 4, 'yokohama_2010_4_ma_q.pdf', 'yokohama_2010_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2019, 3, 'yokohama_2019_3_ma_q.pdf', 'yokohama_2019_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2016, 4, 'yokohama_2016_4_ma_q.pdf', 'yokohama_2016_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2011, 3, 'yokohama_2011_3_ma_q.pdf', 'yokohama_2011_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2010, 7, 'yokohama_2010_7_ma_q.pdf', 'yokohama_2010_7_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2010, 6, 'yokohama_2010_6_ma_q.pdf', 'yokohama_2010_6_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2021, 2, 'yokohama_2021_2_ma_q.pdf', 'yokohama_2021_2_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2021, 3, 'yokohama_2021_3_ma_q.pdf', 'yokohama_2021_3_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2011, 4, 'yokohama_2011_4_ma_q.pdf', 'yokohama_2011_4_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'macro', 2014, 5, 'yokohama_2014_5_ma_q.pdf', 'yokohama_2014_5_ma_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2016, 2, 'yokohama_2016_2_mi_q.pdf', 'yokohama_2016_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2016, 3, 'yokohama_2016_3_mi_q.pdf', 'yokohama_2016_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2020, 1, 'yokohama_2020_1_mi_q.pdf', 'yokohama_2020_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2017, 1, 'yokohama_2017_1_mi_q.pdf', 'yokohama_2017_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2014, 4, 'yokohama_2014_4_mi_q.pdf', 'yokohama_2014_4_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2012, 1, 'yokohama_2012_1_mi_q.pdf', 'yokohama_2012_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2018, 1, 'yokohama_2018_1_mi_q.pdf', 'yokohama_2018_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2014, 3, 'yokohama_2014_3_mi_q.pdf', 'yokohama_2014_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2014, 2, 'yokohama_2014_2_mi_q.pdf', 'yokohama_2014_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2019, 2, 'yokohama_2019_2_mi_q.pdf', 'yokohama_2019_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2015, 1, 'yokohama_2015_1_mi_q.pdf', 'yokohama_2015_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2022, 1, 'yokohama_2022_1_mi_q.pdf', 'yokohama_2022_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2010, 1, 'yokohama_2010_1_mi_q.pdf', 'yokohama_2010_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2011, 2, 'yokohama_2011_2_mi_q.pdf', 'yokohama_2011_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2012, 2, 'yokohama_2012_2_mi_q.pdf', 'yokohama_2012_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2017, 2, 'yokohama_2017_2_mi_q.pdf', 'yokohama_2017_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2010, 5, 'yokohama_2010_5_mi_q.pdf', 'yokohama_2010_5_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2021, 1, 'yokohama_2021_1_mi_q.pdf', 'yokohama_2021_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2016, 1, 'yokohama_2016_1_mi_q.pdf', 'yokohama_2016_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2011, 1, 'yokohama_2011_1_mi_q.pdf', 'yokohama_2011_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2010, 2, 'yokohama_2010_2_mi_q.pdf', 'yokohama_2010_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2010, 3, 'yokohama_2010_3_mi_q.pdf', 'yokohama_2010_3_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2019, 1, 'yokohama_2019_1_mi_q.pdf', 'yokohama_2019_1_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2015, 2, 'yokohama_2015_2_mi_q.pdf', 'yokohama_2015_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2022, 2, 'yokohama_2022_2_mi_q.pdf', 'yokohama_2022_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2018, 2, 'yokohama_2018_2_mi_q.pdf', 'yokohama_2018_2_mi_a.pdf', '', '');
INSERT INTO "exams" ("university", "subject", "year", "question_num", "question_pdf_url", "answer_pdf_url", "video_url", "critique_url") VALUES ('yokohama', 'micro', 2014, 1, 'yokohama_2014_1_mi_q.pdf', 'yokohama_2014_1_mi_a.pdf', '', '');
