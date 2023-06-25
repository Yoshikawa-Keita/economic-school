-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email,
    user_type,
    profile_image_url,
    version
) VALUES (
             $1, $2, $3, $4, $5, $6, 0
         ) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    user_type = COALESCE(sqlc.narg(user_type), user_type),
    profile_image_url = COALESCE(sqlc.narg(profile_image_url), profile_image_url),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified),
    version = version + 1
WHERE
        username = sqlc.arg(username)
    RETURNING *;
