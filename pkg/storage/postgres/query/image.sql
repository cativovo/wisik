-- name: AddImage :one
INSERT INTO image (id, label, src)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: ListImages :many
SELECT * FROM image;
