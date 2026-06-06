package usecase

import (
	"context"
	"testing"
	"time"
)

func TestManualLogin_AdminSuccess(t *testing.T) {
	accounts := &fakeManualAccountRepository{accountID: "acc-admin"}
	roles := &fakeManualRoleAssigner{}
	sessions := &fakeSessionStore{}
	tokens := &fakeTokenIssuer{
		token: "access-token",
		exp:   time.Now().Add(15 * time.Minute),
	}
	tx := &fakeTransactor{}

	usecase := NewManualLogin(accounts, roles, sessions, tokens, tx, 30*24*time.Hour)

	out, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    " ADMIN@example.com ",
		Password: "12345678",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if out.AccessToken != "access-token" {
		t.Fatalf("access token = %q", out.AccessToken)
	}
	if out.RefreshToken == "" {
		t.Fatal("refresh token is empty")
	}
	if out.TrustLevel != "aal1" {
		t.Fatalf("trust level = %q", out.TrustLevel)
	}
	if accounts.lastEmail != "admin@example.com" {
		t.Fatalf("email = %q", accounts.lastEmail)
	}
	if roles.ensureCalls != 1 || roles.lastRoleKey != "admin" {
		t.Fatalf("role assignment = %d/%q, want admin once", roles.ensureCalls, roles.lastRoleKey)
	}
	if sessions.createCalls != 1 {
		t.Fatalf("session create calls = %d, want 1", sessions.createCalls)
	}
	if sessions.lastSession.AccountID != "acc-admin" {
		t.Fatalf("session account id = %q", sessions.lastSession.AccountID)
	}
	if sessions.lastSession.Meta["provider"] != "manual" {
		t.Fatalf("session provider meta = %v", sessions.lastSession.Meta["provider"])
	}
	if tokens.lastReq.AccountID != "acc-admin" {
		t.Fatalf("token account id = %q", tokens.lastReq.AccountID)
	}
	if tokens.lastReq.SessionID != sessions.lastSession.ID {
		t.Fatalf("token session id = %q, want %q", tokens.lastReq.SessionID, sessions.lastSession.ID)
	}
	if tx.runCalls != 1 {
		t.Fatalf("tx run calls = %d, want 1", tx.runCalls)
	}
}
