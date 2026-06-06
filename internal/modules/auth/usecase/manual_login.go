package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

var ErrManualLoginUnsupportedEmail = errors.New("unsupported manual login email")

type ManualLogin struct {
	accounts     ports.ManualAccountRepository
	roles        ports.AccountRoleAssigner
	sessions     ports.SessionStore
	tokens       ports.TokenIssuer
	tx           ports.Transactor
	sessionTTL   time.Duration
	allowedRoles map[string]string
}

func (u *ManualLogin) Execute(ctx context.Context, in ManualLoginInput) (ManualLoginOutput, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))
	roleKey, ok := u.allowedRoles[email]
	if !ok {
		return ManualLoginOutput{}, ErrManualLoginUnsupportedEmail
	}

	seed, err := newManualLoginSeed(u.sessionTTL)
	if err != nil {
		return ManualLoginOutput{}, err
	}

	var accountID string
	if err := u.tx.RunInTx(ctx, func(txCtx context.Context) error {
		resolvedAccountID, err := u.accounts.ResolveOrCreateManualAccount(txCtx, email)
		if err != nil {
			return err
		}
		accountID = resolvedAccountID

		if err := u.roles.EnsureRole(txCtx, accountID, roleKey); err != nil {
			return err
		}

		return u.sessions.Create(txCtx, domain.Session{
			ID:               seed.sessionID,
			AccountID:        accountID,
			RefreshTokenHash: seed.refreshTokenHash,
			CreatedAt:        seed.now,
			ExpiresAt:        seed.refreshExp,
			Meta: map[string]any{
				"provider": "manual",
				"role":     roleKey,
			},
		})
	}); err != nil {
		return ManualLoginOutput{}, err
	}

	accessToken, accessExp, err := u.tokens.IssueAccessToken(ctx, ports.AccessTokenRequest{
		AccountID:  accountID,
		SessionID:  seed.sessionID,
		TrustLevel: "aal1",
	})
	if err != nil {
		return ManualLoginOutput{}, err
	}

	return ManualLoginOutput{
		AccessToken:    accessToken,
		AccessExp:      accessExp,
		RefreshToken:   seed.refreshToken,
		RefreshExp:     seed.refreshExp,
		TrustLevel:     "aal1",
		StepUpRequired: false,
	}, nil
}
