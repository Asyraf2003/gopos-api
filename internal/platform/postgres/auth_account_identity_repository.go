package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountIdentityRepository struct {
	pool *pgxpool.Pool
}

func NewAccountIdentityRepository(pool *pgxpool.Pool) *AccountIdentityRepository {
	return &AccountIdentityRepository{pool: pool}
}

func (r *AccountIdentityRepository) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *AccountIdentityRepository) exec(ctx context.Context, sql string, args ...any) error {
	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, args...)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}

var _ ports.AccountIdentityRepository = (*AccountIdentityRepository)(nil)
