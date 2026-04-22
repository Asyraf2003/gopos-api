package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestRefreshToken_RejectsEmptyRefreshToken(t *testing.T) {
	usecase := NewRefreshToken(&fakeRefreshSessionRepository{}, &fakeTokenIssuer{}, 24*time.Hour)

	_, err := usecase.Execute(context.Background(), RefreshTokenInput{})
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
}

func TestRefreshToken_RejectsExpiredRefreshToken(t *testing.T) {
	repo := &fakeRefreshSessionRepository{
		session: ports.RefreshSession{
			SessionID: "sess-123",
			AccountID: "acc-123",
			ExpiresAt: time.Now().Add(-1 * time.Minute),
		},
	}

	usecase := NewRefreshToken(repo, &fakeTokenIssuer{}, 24*time.Hour)

	_, err := usecase.Execute(context.Background(), RefreshTokenInput{
		RefreshToken: "old-refresh-token",
	})
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
}
