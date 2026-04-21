package domain

import "time"

type Provider string

const ProviderGoogle Provider = "google"

type Identity struct {
	Provider      Provider
	Subject       string
	Email         string
	EmailVerified bool
	Meta          map[string]any
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
