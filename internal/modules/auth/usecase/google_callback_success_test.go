package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
	"pos-go/internal/platform/state/memory"
)

func TestGoogleCallback_Success(t *testing.T) {
	stateStore := memory.NewAuthStateStore()

	oidc := &fakeCallbackOIDCProvider{
		claims: ports.OIDCClaims{
			Provider:      "google",
			Subject:       "google-sub-123",
			Email:         "user@example.com",
			EmailVerified: true,
			AuthTime:      time.Now(),
		},
	}
	accountRepo := &fakeAccountIdentityRepository{
		accountID: "356ef0e8-ea0a-4416-82b6-da91840815d0",
	}
	sessionStore := &fakeSessionStore{}
	tokenIssuer := &fakeTokenIssuer{
		token: "access-token-xyz",
		exp:   time.Now().Add(15 * time.Minute),
	}
	transactor := &fakeTransactor{}
	roleAssigner := &fakeGoogleCallbackRoleAssigner{}

	flow := NewGoogleFlow(
		oidc,
		stateStore,
		accountRepo,
		sessionStore,
		tokenIssuer,
		transactor,
		10*time.Minute,
		24*time.Hour,
	).WithRoleAssigner(roleAssigner)

	err := stateStore.Put(context.Background(), "state-123", ports.AuthState{
		Nonce:        "nonce-123",
		CodeVerifier: "verifier-123",
		Purpose:      "login",
		CreatedAt:    time.Now(),
	}, 10*time.Minute)
	if err != nil {
		t.Fatalf("stateStore.Put() error = %v", err)
	}

	out, err := flow.GoogleCallback(context.Background(), GoogleCallbackInput{
		Code:        "auth-code-123",
		State:       "state-123",
		RedirectURL: "http://127.0.0.1:8081/api/auth/google/callback",
		Client: ClientInfo{
			UserAgent: "test-agent",
			IP:        "127.0.0.1",
		},
	})
	if err != nil {
		t.Fatalf("GoogleCallback() error = %v", err)
	}

	assertGoogleCallbackSuccess(t, out, accountRepo, roleAssigner, sessionStore, tokenIssuer, transactor)

	_, err = stateStore.GetDel(context.Background(), "state-123")
	if err == nil {
		t.Fatal("state should have been consumed")
	}
}
