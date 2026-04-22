package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

func (u *GoogleFlow) resolveGoogleCallbackAccount(
	ctx context.Context,
	claims ports.OIDCClaims,
	now time.Time,
) (string, error) {
	accountID, err := u.accounts.ResolveOrCreateAccountByGoogle(ctx, claims)
	if err != nil {
		return "", err
	}

	if u.roleAssigner != nil {
		if err := u.roleAssigner.EnsureRole(ctx, accountID, "base"); err != nil {
			return "", err
		}
	}

	identity := domain.Identity{
		Provider:      domain.Provider(claims.Provider),
		Subject:       claims.Subject,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		Meta:          map[string]any{},
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := u.accounts.UpsertIdentity(ctx, accountID, identity); err != nil {
		return "", err
	}

	return accountID, nil
}

func (u *GoogleFlow) createGoogleCallbackSession(
	ctx context.Context,
	accountID string,
	seed googleCallbackSeed,
) error {
	session := domain.Session{
		ID:               seed.sessionID,
		AccountID:        accountID,
		RefreshTokenHash: seed.refreshTokenHash,
		CreatedAt:        seed.now,
		ExpiresAt:        seed.refreshExp,
		Meta: map[string]any{
			"provider": "google",
		},
	}

	return u.sessions.Create(ctx, session)
}
