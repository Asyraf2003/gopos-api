package ports

import (
	"context"
	"time"
)

type OIDCAuthURLParams struct {
	State         string
	Nonce         string
	CodeChallenge string
	RedirectURL   string
	Purpose       string
}

type OIDCClaims struct {
	Provider      string
	Subject       string
	Email         string
	EmailVerified bool
	AuthTime      time.Time
}

type OIDCProvider interface {
	AuthCodeURL(p OIDCAuthURLParams) string
	ExchangeAndVerify(ctx context.Context, code, codeVerifier, redirectURL, nonce string) (OIDCClaims, error)
}
