package ports

import "context"

type AccountRoleAssigner interface {
	EnsureRole(ctx context.Context, accountID string, roleKey string) error
}
