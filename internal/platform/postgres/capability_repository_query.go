package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *CapabilityRepository) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}

	return r.pool.Query(ctx, sql, args...)
}

func (r *CapabilityRepository) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}

	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *CapabilityRepository) exec(ctx context.Context, sql string, args ...any) (pgx.CommandTag, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Exec(ctx, sql, args...)
	}

	return r.pool.Exec(ctx, sql, args...)
}
