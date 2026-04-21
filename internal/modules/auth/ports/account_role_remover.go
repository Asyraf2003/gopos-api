package ports

import "context"

type AccountRoleRemover interface {
	RemoveRole(ctx context.Context, accountID string, roleKey string) error
}
