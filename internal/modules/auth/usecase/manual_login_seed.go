package usecase

import (
	"time"

	"github.com/google/uuid"
)

type manualLoginSeed struct {
	now              time.Time
	sessionID        string
	refreshToken     string
	refreshExp       time.Time
	refreshTokenHash string
}

func newManualLoginSeed(sessionTTL time.Duration) (manualLoginSeed, error) {
	now := time.Now()
	refreshToken, err := randB64(48)
	if err != nil {
		return manualLoginSeed{}, err
	}

	return manualLoginSeed{
		now:              now,
		sessionID:        uuid.NewString(),
		refreshToken:     refreshToken,
		refreshExp:       now.Add(sessionTTL),
		refreshTokenHash: sha256Hex(refreshToken),
	}, nil
}
