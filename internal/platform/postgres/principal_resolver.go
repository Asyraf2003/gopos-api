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

func (r *PrincipalResolver) loadRoles(ctx context.Context, accountID string) ([]string, error) {
	rows, err := r.query(ctx, `
		SELECT r.key
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1
		ORDER BY r.key
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleKey string
		if err := rows.Scan(&roleKey); err != nil {
			return nil, err
		}
		roles = append(roles, roleKey)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *PrincipalResolver) loadPermissions(ctx context.Context, accountID string) ([]string, error) {
	rows, err := r.query(ctx, `
		SELECT DISTINCT p.key
		FROM account_roles ar
		JOIN role_permissions rp ON rp.role_id = ar.role_id
		JOIN permissions p ON p.id = rp.permission_id
		WHERE ar.account_id = $1
		ORDER BY p.key
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var permissionKey string
		if err := rows.Scan(&permissionKey); err != nil {
			return nil, err
		}
		permissions = append(permissions, permissionKey)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *PrincipalResolver) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}

	return r.pool.Query(ctx, sql, args...)
}

var _ ports.PrincipalResolver = (*PrincipalResolver)(nil)
