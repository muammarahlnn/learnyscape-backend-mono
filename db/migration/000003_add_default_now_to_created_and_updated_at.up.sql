ALTER TABLE "reset_password_tokens"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();

ALTER TABLE "subjects"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();

ALTER TABLE "subject_classes"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();

ALTER TABLE "classes"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();

ALTER TABLE "class_users"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();

ALTER TABLE "schedules"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();

ALTER TABLE "days"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "announcements"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "modules"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "assignments"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "assignment_submissions"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "quizzes"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "questions"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "choices"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "quiz_attempts"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "quiz_answers"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "attachments"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "attachment_files"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "assignment_grades"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "quiz_grades"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();
                                                
ALTER TABLE "audit_logs"
    ALTER COLUMN "created_at" SET DEFAULT NOW(),
    ALTER COLUMN "updated_at" SET DEFAULT NOW();