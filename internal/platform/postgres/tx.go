package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txContextKey struct{}

type Transactor struct {
	pool *pgxpool.Pool
}

func NewTransactor(pool *pgxpool.Pool) *Transactor {
	return &Transactor{pool: pool}
}

func (t *Transactor) RunInTx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	txCtx := context.WithValue(ctx, txContextKey{}, tx)

	if err := fn(txCtx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	committed = true
	return nil
}

func TxFromContext(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txContextKey{}).(pgx.Tx)
	return tx, ok
}

var _ ports.Transactor = (*Transactor)(nil)
