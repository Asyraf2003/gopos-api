package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/auth/ports"
)

type fakeRefreshSessionRepository struct {
	session          ports.RefreshSession
	findErr          error
	rotateErr        error
	findCalls        int
	rotateCalls      int
	lastLookupHash   string
	lastSessionID    string
	lastNewHash      string
	lastNewExpiresAt time.Time
}

func (f *fakeRefreshSessionRepository) FindActiveByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (ports.RefreshSession, error) {
	_ = ctx
	f.findCalls++
	f.lastLookupHash = refreshTokenHash
	return f.session, f.findErr
}

func (f *fakeRefreshSessionRepository) RotateRefreshToken(ctx context.Context, sessionID string, newRefreshTokenHash string, newExpiresAt time.Time) error {
	_ = ctx
	f.rotateCalls++
	f.lastSessionID = sessionID
	f.lastNewHash = newRefreshTokenHash
	f.lastNewExpiresAt = newExpiresAt
	return f.rotateErr
}
