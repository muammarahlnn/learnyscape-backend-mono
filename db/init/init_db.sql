CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "role_id" bigint NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hash_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "profile_pic_url" varchar
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "reset_password_tokens" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "token" varchar NOT NULL,
  "token_expiry" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "subjects" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "subject_classes" (
  "id" bigserial PRIMARY KEY,
  "subject_id" bigint NOT NULL,
  "class_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "classes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "class_users" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "schedules" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "day_id" bigint NOT NULL,
  "start_time" time NOT NULL,
  "end_time" time NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "days" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "announcements" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "attachment_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "modules" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "attachment_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "assignments" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "attachment_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" text,
  "deadline" timestamptz,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "assignment_submissions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "assignment_id" bigint NOT NULL,
  "attachment_id" bigint NOT NULL,
  "submitted_at" timestamptz,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "quizzes" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" text,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "duration" int NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "questions" (
  "id" bigserial PRIMARY KEY,
  "quiz_id" bigint NOT NULL,
  "question" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "choices" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigint NOT NULL,
  "text" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "quiz_attempts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "quiz_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "quiz_answers" (
  "id" bigserial PRIMARY KEY,
  "quiz_attempt_id" bigint NOT NULL,
  "question_id" bigint NOT NULL,
  "choice_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "attachments" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "attachment_files" (
  "id" bigserial PRIMARY KEY,
  "attachment_id" bigint NOT NULL,
  "url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "assignment_grades" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "assignment_id" bigint NOT NULL,
  "score" decimal NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "quiz_grades" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "quiz_id" bigint NOT NULL,
  "score" decimal NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "audit_logs" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "action" varchar NOT NULL,
  "entity_type" varchar NOT NULL,
  "entity_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "reset_password_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "subject_classes" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "subject_classes" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "class_users" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "class_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "schedules" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "schedules" ADD FOREIGN KEY ("day_id") REFERENCES "days" ("id");

ALTER TABLE "announcements" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "announcements" ADD FOREIGN KEY ("attachment_id") REFERENCES "attachments" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("attachment_id") REFERENCES "attachments" ("id");

ALTER TABLE "assignments" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "assignments" ADD FOREIGN KEY ("attachment_id") REFERENCES "attachments" ("id");

ALTER TABLE "assignment_submissions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "assignment_submissions" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id");

ALTER TABLE "assignment_submissions" ADD FOREIGN KEY ("attachment_id") REFERENCES "attachments" ("id");

ALTER TABLE "quizzes" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "questions" ADD FOREIGN KEY ("quiz_id") REFERENCES "quizzes" ("id");

ALTER TABLE "choices" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "quiz_attempts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "quiz_attempts" ADD FOREIGN KEY ("quiz_id") REFERENCES "quizzes" ("id");

ALTER TABLE "quiz_answers" ADD FOREIGN KEY ("quiz_attempt_id") REFERENCES "quiz_attempts" ("id");

ALTER TABLE "quiz_answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "quiz_answers" ADD FOREIGN KEY ("choice_id") REFERENCES "choices" ("id");

ALTER TABLE "attachment_files" ADD FOREIGN KEY ("attachment_id") REFERENCES "attachments" ("id");

ALTER TABLE "assignment_grades" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "assignment_grades" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id");

ALTER TABLE "quiz_grades" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "quiz_grades" ADD FOREIGN KEY ("quiz_id") REFERENCES "quizzes" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
