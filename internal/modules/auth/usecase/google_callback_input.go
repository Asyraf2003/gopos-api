package usecase

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type googleCallbackSeed struct {
	now              time.Time
	sessionID        string
	refreshToken     string
	refreshExp       time.Time
	refreshTokenHash string
}

func normalizeGoogleCallbackInput(in *GoogleCallbackInput) error {
	in.Code = strings.TrimSpace(in.Code)
	in.State = strings.TrimSpace(in.State)
	in.RedirectURL = strings.TrimSpace(in.RedirectURL)

	if in.Code == "" || in.State == "" || in.RedirectURL == "" {
		return errors.New("code, state, and redirect url are required")
	}

	return nil
}

func newGoogleCallbackSeed(sessionTTL time.Duration) (googleCallbackSeed, error) {
	now := time.Now()

	refreshToken, err := randB64(48)
	if err != nil {
		return googleCallbackSeed{}, err
	}

	return googleCallbackSeed{
		now:              now,
		sessionID:        uuid.NewString(),
		refreshToken:     refreshToken,
		refreshExp:       now.Add(sessionTTL),
		refreshTokenHash: sha256Hex(refreshToken),
	}, nil
}
