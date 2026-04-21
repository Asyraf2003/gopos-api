package ports

import (
	"context"

	"pos-go/internal/modules/auth/domain"
)

type SessionStore interface {
	Create(ctx context.Context, session domain.Session) error
}
