package postgres

import (
	"context"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PrincipalResolver struct {
	pool *pgxpool.Pool
}

func NewPrincipalResolver(pool *pgxpool.Pool) *PrincipalResolver {
	return &PrincipalResolver{pool: pool}
}

func (r *PrincipalResolver) Resolve(ctx context.Context, in ports.ResolvePrincipalInput) (domain.Principal, error) {
	roles, err := r.loadRoles(ctx, in.AccountID)
	if err != nil {
		return domain.Principal{}, err
	}

	permissions, err := r.loadPermissions(ctx, in.AccountID)
	if err != nil {
		return domain.Principal{}, err
	}

	return domain.Principal{
		AccountID:   in.AccountID,
		SessionID:   in.SessionID,
		Roles:       roles,
		Permissions: permissions,
		TrustLevel:  in.TrustLevel,
	}, nil
}

func (r *PrincipalResolver) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}

	return r.pool.Query(ctx, sql, args...)
}

var _ ports.PrincipalResolver = (*PrincipalResolver)(nil)
