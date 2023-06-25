-- name: CreatePasswordResetEmail :one
INSERT INTO password_reset_emails (
    username,
    email,
    secret_code
) VALUES (
             $1, $2, $3
         ) RETURNING *;

-- name: UpdatePasswordResetEmail :one
UPDATE password_reset_emails
SET
    is_used = TRUE
WHERE
        id = @id
  AND secret_code = @secret_code
  AND is_used = FALSE
  AND expired_at > now()
    RETURNING *;
