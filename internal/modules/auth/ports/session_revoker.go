package ports

import "context"

type SessionRevoker interface {
	RevokeSession(ctx context.Context, sessionID string) error
}
