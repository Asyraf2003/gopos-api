package ports

import (
	"context"
	"time"
)

type RefreshSession struct {
	SessionID        string
	AccountID        string
	RefreshTokenHash string
	ExpiresAt        time.Time
	RevokedAt        *time.Time
}

type RefreshSessionRepository interface {
	FindActiveByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (RefreshSession, error)
	RotateRefreshToken(ctx context.Context, sessionID string, newRefreshTokenHash string, newExpiresAt time.Time) error
}
