CREATE TABLE "shorturls" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer NOT NUll,
  "origin" varchar NOT NULL,
  "match" varchar(20) UNIQUE NOT NULL,
  "expired_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "line_id" varchar(100) UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "exchage_tokens" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer,
  "token" varchar(100) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE INDEX ON "shorturls" ("user_id");

CREATE INDEX ON "shorturls" ("match", "expired_at");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("line_id");

CREATE INDEX ON "exchage_tokens" ("token");

ALTER TABLE "exchage_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
