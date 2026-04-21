package ports

import (
	"context"

	"pos-go/internal/modules/auth/domain"
)

type AccountIdentityRepository interface {
	ResolveOrCreateAccountByGoogle(ctx context.Context, claims OIDCClaims) (accountID string, err error)
	UpsertIdentity(ctx context.Context, accountID string, identity domain.Identity) error
}
