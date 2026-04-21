package ports

import (
	"context"

	"pos-go/internal/modules/auth/domain"
)

type ResolvePrincipalInput struct {
	AccountID  string
	SessionID  string
	TrustLevel string
}

type PrincipalResolver interface {
	Resolve(ctx context.Context, in ResolvePrincipalInput) (domain.Principal, error)
}
