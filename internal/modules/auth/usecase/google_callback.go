package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/google/uuid"
)

func (u *GoogleFlow) GoogleCallback(ctx context.Context, in GoogleCallbackInput) (GoogleCallbackOutput, error) {
	in.Code = strings.TrimSpace(in.Code)
	in.State = strings.TrimSpace(in.State)
	in.RedirectURL = strings.TrimSpace(in.RedirectURL)

	if in.Code == "" || in.State == "" || in.RedirectURL == "" {
		return GoogleCallbackOutput{}, errors.New("code, state, and redirect url are required")
	}

	storedState, err := u.states.GetDel(ctx, in.State)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	claims, err := u.oidc.ExchangeAndVerify(
		ctx,
		in.Code,
		storedState.CodeVerifier,
		in.RedirectURL,
		storedState.Nonce,
	)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	now := time.Now()

	sessionID := uuid.NewString()

	refreshToken, err := randB64(48)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	refreshExp := now.Add(u.sessionTTL)
	refreshTokenHash := sha256Hex(refreshToken)

	var accountID string

	err = u.tx.RunInTx(ctx, func(txCtx context.Context) error {
		resolvedAccountID, err := u.accounts.ResolveOrCreateAccountByGoogle(txCtx, claims)
		if err != nil {
			return err
		}

		accountID = resolvedAccountID

		if u.roleAssigner != nil {
			if err := u.roleAssigner.EnsureRole(txCtx, accountID, "base"); err != nil {
				return err
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

		if err := u.accounts.UpsertIdentity(txCtx, accountID, identity); err != nil {
			return err
		}

		session := domain.Session{
			ID:               sessionID,
			AccountID:        accountID,
			RefreshTokenHash: refreshTokenHash,
			CreatedAt:        now,
			ExpiresAt:        refreshExp,
			Meta: map[string]any{
				"provider": "google",
			},
		}

		return u.sessions.Create(txCtx, session)
	})
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	accessToken, accessExp, err := u.tokens.IssueAccessToken(ctx, ports.AccessTokenRequest{
		AccountID:  accountID,
		SessionID:  sessionID,
		TrustLevel: "aal1",
	})
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	return GoogleCallbackOutput{
		AccessToken:    accessToken,
		AccessExp:      accessExp,
		RefreshToken:   refreshToken,
		RefreshExp:     refreshExp,
		TrustLevel:     "aal1",
		StepUpRequired: false,
	}, nil
}
