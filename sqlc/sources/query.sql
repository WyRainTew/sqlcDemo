-- name: GetBook :one
SELECT * FROM books WHERE id = $1;

-- name: GetAllBooks :many
SELECT * FROM books ORDER BY name;

-- name: GetBooksByCategory :many
SELECT * FROM books WHERE category = $1 ORDER BY name;