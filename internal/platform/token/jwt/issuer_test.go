package jwt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestIssueAccessToken_EncodesExpectedClaims(t *testing.T) {
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

	token, exp, err := issuer.IssueAccessToken(context.Background(), ports.AccessTokenRequest{
		AccountID:  "356ef0e8-ea0a-4416-82b6-da91840815d0",
		SessionID:  "fce0c7d0-903f-4bdf-82c8-393d1c292b48",
		TrustLevel: "aal1",
	})
	if err != nil {
		t.Fatalf("IssueAccessToken() error = %v", err)
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("token parts = %d, want 3", len(parts))
	}

	payloadJSON := decodeBase64URL(t, parts[1])

	var payload map[string]any
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if got := payload["iss"]; got != "pos-go" {
		t.Fatalf("iss = %v, want pos-go", got)
	}
	if got := payload["aud"]; got != "pos-go-client" {
		t.Fatalf("aud = %v, want pos-go-client", got)
	}
	if got := payload["sub"]; got != "356ef0e8-ea0a-4416-82b6-da91840815d0" {
		t.Fatalf("sub = %v", got)
	}
	if got := payload["sid"]; got != "fce0c7d0-903f-4bdf-82c8-393d1c292b48" {
		t.Fatalf("sid = %v", got)
	}
	if got := payload["aal"]; got != "aal1" {
		t.Fatalf("aal = %v, want aal1", got)
	}

	iat, ok := payload["iat"].(float64)
	if !ok {
		t.Fatalf("iat missing or invalid type: %T", payload["iat"])
	}

	expClaim, ok := payload["exp"].(float64)
	if !ok {
		t.Fatalf("exp missing or invalid type: %T", payload["exp"])
	}

	if int64(expClaim) <= int64(iat) {
		t.Fatalf("exp <= iat: exp=%v iat=%v", expClaim, iat)
	}

	if exp.IsZero() {
		t.Fatal("returned exp is zero")
	}
}

func decodeBase64URL(t *testing.T, value string) []byte {
	t.Helper()

	raw, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		t.Fatalf("DecodeString() error = %v", err)
	}

	return raw
}
