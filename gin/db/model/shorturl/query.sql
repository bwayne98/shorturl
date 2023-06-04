CREATE TABLE "shorturls" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer,
  "origin" varchar NOT NULL,
  "match" varchar(20) UNIQUE NOT NULL,
  "expired_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

-- name: CreateShorturl :one
INSERT INTO shorturls (
  origin, match, user_id, expired_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetMatchShorturl :one
SELECT origin FROM shorturls 
WHERE match = $1 AND expired_at > (now()) 
ORDER BY id LIMIT 1;

-- name: UpdateExpired :one
UPDATE shorturls 
SET expired_at = $2 
WHERE id = $1 
RETURNING *;

-- name: DeleteShorturl :exec
DELETE FROM shorturls 
WHERE id = $1 
AND user_id = $2;

-- name: ListUserShorturl :many
SELECT id, origin, match, expired_at, created_at
FROM shorturls
WHERE user_id = $1;

-- name: CountMatchShorturl :one
SELECT COUNT(id) FROM shorturls WHERE match = $1;