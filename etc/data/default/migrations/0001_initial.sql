-- +goose Up
CREATE TABLE IF NOT EXISTS "user" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "kind" TEXT NOT NULL,
  "name" TEXT NOT NULL,
  "email" TEXT CONSTRAINT "user_email_unique" UNIQUE,
  "phone" TEXT,
  "github_handle" TEXT,
  "x_handle" TEXT,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  "github_remote_id" TEXT CONSTRAINT "user_github_remote_id_unique" UNIQUE,
  "x_remote_id" TEXT,
  "individual_profile_id" CHAR(26)
);

CREATE TABLE IF NOT EXISTS "profile" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "kind" TEXT NOT NULL,
  "slug" TEXT NOT NULL CONSTRAINT "profile_slug_unique" UNIQUE,
  "profile_picture_uri" TEXT,
  "title" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "show_stories" BOOLEAN DEFAULT FALSE NOT NULL,
  "show_projects" BOOLEAN DEFAULT FALSE NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS "profile_membership" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "kind" TEXT NOT NULL,
  "profile_id" CHAR(26) NOT NULL,
  "user_id" CHAR(26) NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  CONSTRAINT "profile_membership_profile_id_user_id_unique" UNIQUE ("profile_id", "user_id")
);

CREATE TABLE IF NOT EXISTS "session" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "status" TEXT NOT NULL,
  "oauth_request_state" TEXT NOT NULL,
  "oauth_request_code_verifier" TEXT NOT NULL,
  "oauth_redirect_uri" TEXT,
  "logged_in_user_id" CHAR(26),
  "logged_in_at" TIMESTAMP WITH TIME ZONE,
  "expires_at" TIMESTAMP WITH TIME ZONE,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS "session_logged_in_user_id_index" ON "session" ("logged_in_user_id");

CREATE TABLE IF NOT EXISTS "question" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "user_id" CHAR(26) NOT NULL,
  "content" TEXT NOT NULL,
  "is_hidden" BOOLEAN DEFAULT FALSE NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  "answered_at" TIMESTAMP WITH TIME ZONE,
  "answer_uri" TEXT,
  "is_anonymous" BOOLEAN DEFAULT FALSE NOT NULL,
  "answer_kind" TEXT,
  "answer_content" TEXT
);

CREATE TABLE IF NOT EXISTS "question_vote" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "question_id" CHAR(26) NOT NULL,
  "user_id" CHAR(26) NOT NULL,
  "score" INTEGER NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  CONSTRAINT "question_vote_question_id_user_id_unique" UNIQUE ("question_id", "user_id")
);

CREATE TABLE IF NOT EXISTS "event" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "kind" TEXT NOT NULL,
  "slug" TEXT NOT NULL CONSTRAINT "event_slug_unique" UNIQUE,
  "event_picture_uri" TEXT,
  "title" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "time_start" TIMESTAMP WITH TIME ZONE NOT NULL,
  "time_end" TIMESTAMP WITH TIME ZONE NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  "series_id" CHAR(26),
  "status" TEXT DEFAULT 'draft'::TEXT NOT NULL,
  "attendance_uri" TEXT,
  "published_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS "event_attendance" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "kind" TEXT NOT NULL,
  "event_id" CHAR(26) NOT NULL,
  "profile_id" CHAR(26) NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  CONSTRAINT "event_attendance_event_id_profile_id_unique" UNIQUE ("event_id", "profile_id")
);

CREATE TABLE IF NOT EXISTS "event_series" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "slug" TEXT NOT NULL CONSTRAINT "event_series_slug_unique" UNIQUE,
  "event_picture_uri" TEXT,
  "title" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS "story" (
  "id" CHAR(26) NOT NULL PRIMARY KEY,
  "kind" TEXT NOT NULL,
  "status" TEXT NOT NULL,
  "is_featured" BOOLEAN DEFAULT FALSE,
  "slug" TEXT NOT NULL CONSTRAINT "story_slug_unique" UNIQUE,
  "story_picture_uri" TEXT,
  "title" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "author_profile_id" CHAR(26),
  "content" TEXT NOT NULL,
  "published_at" TIMESTAMP WITH TIME ZONE,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  "summary" TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "story";

DROP TABLE IF EXISTS "event_series";

DROP TABLE IF EXISTS "event_attendance";

DROP TABLE IF EXISTS "event";

DROP TABLE IF EXISTS "question_vote";

DROP TABLE IF EXISTS "question";

DROP INDEX IF EXISTS "session_logged_in_user_id_index";

DROP TABLE IF EXISTS "session";

DROP TABLE IF EXISTS "profile_membership";

DROP TABLE IF EXISTS "profile";

DROP TABLE IF EXISTS "user";
