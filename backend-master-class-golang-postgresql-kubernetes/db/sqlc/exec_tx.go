package db

import (
	"context"
	"fmt"
)

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx) // 分離レベルはデフォルト
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		// ロールバック
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx) // コミット
}
