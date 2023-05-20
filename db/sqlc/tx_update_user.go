package db

import "context"

type UpdateUserTxParams struct {
	UpdateUserParams
	UpdateImage func(user User) error
}

type UpdateUserTxResult struct {
	User User
}

func (store *SQLStore) UpdateUserTx(ctx context.Context, arg UpdateUserTxParams) (UpdateUserTxResult, error) {
	var result UpdateUserTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.User, err = q.UpdateUser(ctx, arg.UpdateUserParams)
		if err != nil {
			return err
		}
		uploadErr := arg.UpdateImage(result.User)
		if uploadErr != nil {
			return uploadErr
		}

		return nil
	})

	return result, err
}
