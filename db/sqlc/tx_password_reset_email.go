package db

import (
	"context"
	"database/sql"
)

type PasswordResetEmailTxParams struct {
	EmailId        int64
	SecretCode     string
	HashedPassword string
}

type PasswordResetEmailTxResult struct {
	User               User
	PasswordResetEmail PasswordResetEmail
}

func (store *SQLStore) PasswordResetEmailTx(ctx context.Context, arg PasswordResetEmailTxParams) (PasswordResetEmailTxResult, error) {
	var result PasswordResetEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.PasswordResetEmail, err = q.UpdatePasswordResetEmail(ctx, UpdatePasswordResetEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.PasswordResetEmail.Username,
			HashedPassword: sql.NullString{
				String: arg.HashedPassword,
				Valid:  true,
			},
		})
		return err
	})

	return result, err
}
