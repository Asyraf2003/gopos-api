package usecase

import (
	"context"
	"strings"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
	"pos-go/internal/platform/state/memory"
)

type fakeOIDCProvider struct{}

func (f *fakeOIDCProvider) AuthCodeURL(p ports.OIDCAuthURLParams) string {
	return "https://example.com/oauth?state=" + p.State + "&nonce=" + p.Nonce + "&redirect=" + p.RedirectURL + "&purpose=" + p.Purpose
}

func (f *fakeOIDCProvider) ExchangeAndVerify(ctx context.Context, code, codeVerifier, redirectURL, nonce string) (ports.OIDCClaims, error) {
	_ = ctx
	_ = code
	_ = codeVerifier
	_ = redirectURL
	_ = nonce
	return ports.OIDCClaims{}, nil
}

func TestGoogleStart_DefaultsToLoginAndStoresState(t *testing.T) {
	stateStore := memory.NewAuthStateStore()

	flow := NewGoogleFlow(
		&fakeOIDCProvider{},
		stateStore,
		nil,
		nil,
		nil,
		nil,
		10*time.Minute,
		24*time.Hour,
	)

	out, err := flow.GoogleStart(context.Background(), GoogleStartInput{
		RedirectURL: "http://127.0.0.1:8081/api/auth/google/callback",
	})
	if err != nil {
		t.Fatalf("GoogleStart() error = %v", err)
	}

	if strings.TrimSpace(out.State) == "" {
		t.Fatal("GoogleStart() returned empty state")
	}

	if !strings.Contains(out.RedirectTo, "https://example.com/oauth?") {
		t.Fatalf("redirect_to = %q, want fake oidc url", out.RedirectTo)
	}

	stored, err := stateStore.GetDel(context.Background(), out.State)
	if err != nil {
		t.Fatalf("stateStore.GetDel() error = %v", err)
	}

	if stored.Purpose != "login" {
		t.Fatalf("stored purpose = %q, want login", stored.Purpose)
	}

	if stored.Nonce == "" {
		t.Fatal("stored nonce is empty")
	}

	if stored.CodeVerifier == "" {
		t.Fatal("stored code verifier is empty")
	}
}

func TestGoogleStart_RejectsEmptyRedirectURL(t *testing.T) {
	stateStore := memory.NewAuthStateStore()

	flow := NewGoogleFlow(
		&fakeOIDCProvider{},
		stateStore,
		nil,
		nil,
		nil,
		nil,
		10*time.Minute,
		24*time.Hour,
	)

	_, err := flow.GoogleStart(context.Background(), GoogleStartInput{})
	if err == nil {
		t.Fatal("GoogleStart() error = nil, want error")
	}
}
