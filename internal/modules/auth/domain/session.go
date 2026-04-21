package domain

import "time"

type Session struct {
	ID               string
	AccountID        string
	RefreshTokenHash string
	CreatedAt        time.Time
	ExpiresAt        time.Time
	RevokedAt        *time.Time
	Meta             map[string]any
}
