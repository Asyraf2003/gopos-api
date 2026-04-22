package jwt

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func mustIssueTestAccessToken(
	t *testing.T,
	ttl time.Duration,
	secret string,
	accountID string,
	sessionID string,
	trustLevel string,
) string {
	t.Helper()

	issuer, err := NewHMACIssuer(
		"pos-go",
		"pos-go-client",
		"local-dev-key",
		secret,
		ttl,
	)
	if err != nil {
		t.Fatalf("NewHMACIssuer() error = %v", err)
	}

	token, _, err := issuer.IssueAccessToken(context.Background(), ports.AccessTokenRequest{
		AccountID:  accountID,
		SessionID:  sessionID,
		TrustLevel: trustLevel,
	})
	if err != nil {
		t.Fatalf("IssueAccessToken() error = %v", err)
	}

	return token
}

func mustNewTestVerifier(t *testing.T, secret string) *Verifier {
	t.Helper()

	verifier, err := NewHMACVerifier("pos-go", "pos-go-client", secret)
	if err != nil {
		t.Fatalf("NewHMACVerifier() error = %v", err)
	}

	return verifier
}
