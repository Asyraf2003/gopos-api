package usecase

import (
	"context"
	"errors"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

type LogoutCurrentSession struct {
	revoker ports.SessionRevoker
}

func NewLogoutCurrentSession(revoker ports.SessionRevoker) *LogoutCurrentSession {
	return &LogoutCurrentSession{revoker: revoker}
}

func (u *LogoutCurrentSession) Execute(ctx context.Context, principal domain.Principal) error {
	if principal.SessionID == "" {
		return errors.New("session id is required")
	}

	return u.revoker.RevokeSession(ctx, principal.SessionID)
}
