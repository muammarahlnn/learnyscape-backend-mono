CREATE TABLE IF NOT EXISTS "user_verifications" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "token" varchar NOT NULL,
  "expire_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "user_verifications"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
