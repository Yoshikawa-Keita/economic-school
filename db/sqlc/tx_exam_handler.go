package db

import "context"

type CreateExamTxParams struct {
	CreateExamParams
}

type CreateExamTxResult struct {
	Exam
}

func (store *SQLStore) CreateExamTx(ctx context.Context, arg CreateExamTxParams) (CreateExamTxResult, error) {
	var result CreateExamTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Exam, err = q.CreateExam(ctx, arg.CreateExamParams)
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
