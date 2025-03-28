-- name: GetProfileById :one
SELECT * FROM "profile"
WHERE id = $1
LIMIT 1;

-- name: GetProfileBySlug :one
SELECT * FROM "profile"
WHERE slug = $1
LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM "profile";

-- name: CreateProfile :one
INSERT INTO "profile" (id, slug)
VALUES ($1, $2) RETURNING *;

-- name: UpdateProfile :execrows
UPDATE "profile"
SET slug = $2
WHERE id = $1;

-- name: DeleteProfile :execrows
UPDATE "profile"
SET deleted_at = NOW()
WHERE id = $1;
