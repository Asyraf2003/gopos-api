package usecase

import (
	"context"
	"strings"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestRefreshToken_Success(t *testing.T) {
	repo := &fakeRefreshSessionRepository{
		session: ports.RefreshSession{
			SessionID:        "sess-123",
			AccountID:        "acc-123",
			RefreshTokenHash: sha256Hex("old-refresh-token"),
			ExpiresAt:        time.Now().Add(24 * time.Hour),
			RevokedAt:        nil,
		},
	}

	tokenIssuer := &fakeTokenIssuer{
		token: "new-access-token",
		exp:   time.Now().Add(15 * time.Minute),
	}

	usecase := NewRefreshToken(repo, tokenIssuer, 30*24*time.Hour)

	out, err := usecase.Execute(context.Background(), RefreshTokenInput{
		RefreshToken: "old-refresh-token",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if out.AccessToken != "new-access-token" {
		t.Fatalf("access token = %q", out.AccessToken)
	}
	if strings.TrimSpace(out.RefreshToken) == "" {
		t.Fatal("refresh token is empty")
	}
	if out.RefreshToken == "old-refresh-token" {
		t.Fatal("refresh token was not rotated")
	}
	if out.TrustLevel != "aal1" {
		t.Fatalf("trust level = %q", out.TrustLevel)
	}
	if out.StepUpRequired {
		t.Fatal("step_up_required = true, want false")
	}

	if repo.findCalls != 1 {
		t.Fatalf("find calls = %d, want 1", repo.findCalls)
	}
	if repo.lastLookupHash != sha256Hex("old-refresh-token") {
		t.Fatalf("lookup hash = %q", repo.lastLookupHash)
	}
	if repo.rotateCalls != 1 {
		t.Fatalf("rotate calls = %d, want 1", repo.rotateCalls)
	}
	if repo.lastSessionID != "sess-123" {
		t.Fatalf("rotated session id = %q", repo.lastSessionID)
	}
	if repo.lastNewHash == "" {
		t.Fatal("new refresh token hash is empty")
	}
	if repo.lastNewHash == sha256Hex("old-refresh-token") {
		t.Fatal("new refresh token hash was not rotated")
	}

	if tokenIssuer.issueCalls != 1 {
		t.Fatalf("token issue calls = %d, want 1", tokenIssuer.issueCalls)
	}
	if tokenIssuer.lastReq.AccountID != "acc-123" {
		t.Fatalf("token issuer account id = %q", tokenIssuer.lastReq.AccountID)
	}
	if tokenIssuer.lastReq.SessionID != "sess-123" {
		t.Fatalf("token issuer session id = %q", tokenIssuer.lastReq.SessionID)
	}
	if tokenIssuer.lastReq.TrustLevel != "aal1" {
		t.Fatalf("token issuer trust level = %q", tokenIssuer.lastReq.TrustLevel)
	}
}
