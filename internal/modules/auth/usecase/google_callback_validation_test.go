package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/platform/state/memory"
)

func TestGoogleCallback_RejectsMissingFields(t *testing.T) {
	flow := NewGoogleFlow(
		&fakeCallbackOIDCProvider{},
		memory.NewAuthStateStore(),
		nil,
		nil,
		nil,
		nil,
		10*time.Minute,
		24*time.Hour,
	)

	_, err := flow.GoogleCallback(context.Background(), GoogleCallbackInput{})
	if err == nil {
		t.Fatal("GoogleCallback() error = nil, want error")
	}
}
