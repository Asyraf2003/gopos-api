package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/auth/ports"
)

var ErrBaseRoleRemovalNotAllowed = errors.New("base role cannot be removed")

type RemoveAccountRole struct {
	remover ports.AccountRoleRemover
}

func NewRemoveAccountRole(remover ports.AccountRoleRemover) *RemoveAccountRole {
	return &RemoveAccountRole{remover: remover}
}

func (u *RemoveAccountRole) Execute(ctx context.Context, accountID string, roleKey string) error {
	accountID = strings.TrimSpace(accountID)
	roleKey = strings.TrimSpace(roleKey)

	if accountID == "" {
		return errors.New("account id is required")
	}
	if roleKey == "" {
		return errors.New("role key is required")
	}
	if roleKey == "base" {
		return ErrBaseRoleRemovalNotAllowed
	}

	return u.remover.RemoveRole(ctx, accountID, roleKey)
}
