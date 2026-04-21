package ports

import "context"

type AccessTokenClaims struct {
	AccountID  string
	SessionID  string
	TrustLevel string
}

type AccessTokenVerifier interface {
	VerifyAccessToken(ctx context.Context, token string) (AccessTokenClaims, error)
}
