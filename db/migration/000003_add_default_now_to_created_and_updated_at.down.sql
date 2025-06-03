ALTER TABLE "reset_password_tokens"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;

ALTER TABLE "subjects"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;

ALTER TABLE "subject_classes"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;

ALTER TABLE "classes"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;

ALTER TABLE "class_users"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;

ALTER TABLE "schedules"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;

ALTER TABLE "days"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "announcements"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "modules"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "assignments"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "assignment_submissions"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "quizzes"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "questions"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "choices"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "quiz_attempts"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "quiz_answers"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "attachments"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "attachment_files"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "assignment_grades"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "quiz_grades"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;
                                                
ALTER TABLE "audit_logs"
    ALTER COLUMN "created_at" DROP DEFAULT,
    ALTER COLUMN "updated_at" DROP DEFAULT;