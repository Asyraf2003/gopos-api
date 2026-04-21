package usecase

import (
	"time"

	"pos-go/internal/modules/auth/ports"
)

type GoogleFlow struct {
	oidc         ports.OIDCProvider
	states       ports.AuthStateStore
	accounts     ports.AccountIdentityRepository
	sessions     ports.SessionStore
	tokens       ports.TokenIssuer
	tx           ports.Transactor
	roleAssigner ports.AccountRoleAssigner
	stateTTL     time.Duration
	sessionTTL   time.Duration
}

func NewGoogleFlow(
	oidc ports.OIDCProvider,
	states ports.AuthStateStore,
	accounts ports.AccountIdentityRepository,
	sessions ports.SessionStore,
	tokens ports.TokenIssuer,
	tx ports.Transactor,
	stateTTL time.Duration,
	sessionTTL time.Duration,
) *GoogleFlow {
	return &GoogleFlow{
		oidc:       oidc,
		states:     states,
		accounts:   accounts,
		sessions:   sessions,
		tokens:     tokens,
		tx:         tx,
		stateTTL:   stateTTL,
		sessionTTL: sessionTTL,
	}
}

func (u *GoogleFlow) WithRoleAssigner(roleAssigner ports.AccountRoleAssigner) *GoogleFlow {
	u.roleAssigner = roleAssigner
	return u
}
