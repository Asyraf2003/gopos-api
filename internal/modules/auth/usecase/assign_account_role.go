package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/auth/ports"
)

type AssignAccountRole struct {
	assigner ports.AccountRoleAssigner
}

func NewAssignAccountRole(assigner ports.AccountRoleAssigner) *AssignAccountRole {
	return &AssignAccountRole{assigner: assigner}
}

func (u *AssignAccountRole) Execute(ctx context.Context, accountID string, roleKey string) error {
	accountID = strings.TrimSpace(accountID)
	roleKey = strings.TrimSpace(roleKey)

	if accountID == "" {
		return errors.New("account id is required")
	}
	if roleKey == "" {
		return errors.New("role key is required")
	}

	return u.assigner.EnsureRole(ctx, accountID, roleKey)
}
