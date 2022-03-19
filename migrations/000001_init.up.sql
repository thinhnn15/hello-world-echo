CREATE TABLE "users" (
                         "user_id" text PRIMARY KEY,
                         "full_name" text,
                         "email" text UNIQUE,
                         "password" text,
                         "role" text,
                         "created_at" timestamptz NOT NULL,
                         "updated_at" timestamptz NOT NULL
);

CREATE TABLE "repos" (
                         "name" text PRIMARY KEY,
                         "description" text,
                         "url" text,
                         "color" text,
                         "lang" text,
                         "fork" text,
                         "stars" text,
                         "start_today" text,
                         "build_by" text,
                         "created_at" timestamptz NOT NULL,
                         "updated_at" timestamptz NOT NULL
);

CREATE TABLE "bookmarks" (
                             "bid" text PRIMARY KEY,
                             "user_id" text,
                             "repo_name" text,
                             "created_at" timestamptz NOT NULL,
                             "updated_at" timestamptz NOT NULL
);

ALTER TABLE "bookmarks"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "bookmarks"
    ADD FOREIGN KEY ("repo_name") REFERENCES "repos" ("name");