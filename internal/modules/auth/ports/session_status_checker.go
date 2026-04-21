package ports

import "context"

type SessionStatusChecker interface {
	IsSessionActive(ctx context.Context, sessionID string) (bool, error)
}
