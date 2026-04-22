package jwt

import (
	"context"
	"testing"
	"time"
)

func TestVerifierVerifyAccessToken_Success(t *testing.T) {
	token := mustIssueTestAccessToken(
		t,
		15*time.Minute,
		"test-secret-123",
		"356ef0e8-ea0a-4416-82b6-da91840815d0",
		"fce0c7d0-903f-4bdf-82c8-393d1c292b48",
		"aal1",
	)

	verifier := mustNewTestVerifier(t, "test-secret-123")

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
	token := mustIssueTestAccessToken(
		t,
		15*time.Minute,
		"test-secret-123",
		"acc-1",
		"sess-1",
		"aal1",
	)

	verifier := mustNewTestVerifier(t, "wrong-secret")

	_, err := verifier.VerifyAccessToken(context.Background(), token)
	if err == nil {
		t.Fatal("VerifyAccessToken() error = nil, want error")
	}
}

func TestVerifierVerifyAccessToken_RejectsExpiredToken(t *testing.T) {
	token := mustIssueTestAccessToken(
		t,
		1*time.Minute,
		"test-secret-123",
		"acc-1",
		"sess-1",
		"aal1",
	)

	verifier := mustNewTestVerifier(t, "test-secret-123")
	verifier.nowFn = func() time.Time {
		return time.Now().Add(2 * time.Minute)
	}

	_, err := verifier.VerifyAccessToken(context.Background(), token)
	if err == nil {
		t.Fatal("VerifyAccessToken() error = nil, want error")
	}
}
