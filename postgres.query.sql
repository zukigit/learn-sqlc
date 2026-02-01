-- name: GetWriters :one
SELECT * FROM writers
WHERE id = $1 LIMIT 1;

-- name: ListWriterss :many
SELECT * FROM writers
ORDER BY name;

-- name: CreateWriters :one
INSERT INTO writers (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateWriters :exec
UPDATE writers
  set name = $2,
  bio = $3
WHERE id = $1;

-- name: DeleteWriters :exec
DELETE FROM writers
WHERE id = $1;

-- name: DeleteWritersWithName :exec
DELETE FROM writers
WHERE name = $1;

-- name: CountWriterss :one
SELECT count(*) FROM writers;

-- name: CreateBooks :one
INSERT INTO books (
  writer_id, name, bio
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteBooks :exec
DELETE FROM books
WHERE id = $1;

-- name: WritersAndBooks :many
SELECT sqlc.embed(writers), sqlc.embed(books)
FROM books
JOIN writers ON writers.id = books.writer_id
WHERE books.id = $1;