package jwt

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestVerifierVerifyAccessToken_Success(t *testing.T) {
	issuer, err := NewHMACIssuer(
		"pos-go",
		"pos-go-client",
		"local-dev-key",
		"test-secret-123",
		15*time.Minute,
	)
	if err != nil {
		t.Fatalf("NewHMACIssuer() error = %v", err)
	}

	token, _, err := issuer.IssueAccessToken(context.Background(), ports.AccessTokenRequest{
		AccountID:  "356ef0e8-ea0a-4416-82b6-da91840815d0",
		SessionID:  "fce0c7d0-903f-4bdf-82c8-393d1c292b48",
		TrustLevel: "aal1",
	})
	if err != nil {
		t.Fatalf("IssueAccessToken() error = %v", err)
	}

	verifier, err := NewHMACVerifier("pos-go", "pos-go-client", "test-secret-123")
	if err != nil {
		t.Fatalf("NewHMACVerifier() error = %v", err)
	}

	claims, err := verifier.VerifyAccessToken(context.Background(), token)
	if err != nil {
		t.Fatalf("VerifyAccessToken() error = %v", err)
	}

	if claims.AccountID != "356ef0e8-ea0a-4416-82b6-da91840815d0" {
		t.Fatalf("account id = %q", claims.AccountID)
	}
	if claims.SessionID != "fce0c7d0-903f-4bdf-82c8-393d1c292b48" {
		t.Fatalf("session id = %q", claims.SessionID)
	}
	if claims.TrustLevel != "aal1" {
		t.Fatalf("trust level = %q", claims.TrustLevel)
	}
}

func TestVerifierVerifyAccessToken_RejectsWrongSecret(t *testing.T) {
	issuer, err := NewHMACIssuer(
		"pos-go",
		"pos-go-client",
		"local-dev-key",
		"test-secret-123",
		15*time.Minute,
	)
	if err != nil {
		t.Fatalf("NewHMACIssuer() error = %v", err)
	}

	token, _, err := issuer.IssueAccessToken(context.Background(), ports.AccessTokenRequest{
		AccountID:  "acc-1",
		SessionID:  "sess-1",
		TrustLevel: "aal1",
	})
	if err != nil {
		t.Fatalf("IssueAccessToken() error = %v", err)
	}

	verifier, err := NewHMACVerifier("pos-go", "pos-go-client", "wrong-secret")
	if err != nil {
		t.Fatalf("NewHMACVerifier() error = %v", err)
	}

	_, err = verifier.VerifyAccessToken(context.Background(), token)
	if err == nil {
		t.Fatal("VerifyAccessToken() error = nil, want error")
	}
}

func TestVerifierVerifyAccessToken_RejectsExpiredToken(t *testing.T) {
	issuer, err := NewHMACIssuer(
		"pos-go",
		"pos-go-client",
		"local-dev-key",
		"test-secret-123",
		1*time.Minute,
	)
	if err != nil {
		t.Fatalf("NewHMACIssuer() error = %v", err)
	}

	token, _, err := issuer.IssueAccessToken(context.Background(), ports.AccessTokenRequest{
		AccountID:  "acc-1",
		SessionID:  "sess-1",
		TrustLevel: "aal1",
	})
	if err != nil {
		t.Fatalf("IssueAccessToken() error = %v", err)
	}

	verifier, err := NewHMACVerifier("pos-go", "pos-go-client", "test-secret-123")
	if err != nil {
		t.Fatalf("NewHMACVerifier() error = %v", err)
	}
	verifier.nowFn = func() time.Time {
		return time.Now().Add(2 * time.Minute)
	}

	_, err = verifier.VerifyAccessToken(context.Background(), token)
	if err == nil {
		t.Fatal("VerifyAccessToken() error = nil, want error")
	}
}
