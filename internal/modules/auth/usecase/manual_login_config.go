package usecase

import (
	"time"

	"pos-go/internal/modules/auth/ports"
)

func NewManualLogin(
	accounts ports.ManualAccountRepository,
	roles ports.AccountRoleAssigner,
	sessions ports.SessionStore,
	tokens ports.TokenIssuer,
	tx ports.Transactor,
	sessionTTL time.Duration,
) *ManualLogin {
	return &ManualLogin{
		accounts:   accounts,
		roles:      roles,
		sessions:   sessions,
		tokens:     tokens,
		tx:         tx,
		sessionTTL: sessionTTL,
		allowedRoles: map[string]string{
			"admin@example.com": "admin",
			"kasir@example.com": "cashier",
		},
	}
}
