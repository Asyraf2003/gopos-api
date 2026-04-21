package usecase

import (
	"strings"
	"testing"
)

func assertGoogleCallbackSuccess(
	t *testing.T,
	out GoogleCallbackOutput,
	accountRepo *fakeAccountIdentityRepository,
	roleAssigner *fakeGoogleCallbackRoleAssigner,
	sessionStore *fakeSessionStore,
	tokenIssuer *fakeTokenIssuer,
	transactor *fakeTransactor,
) {
	t.Helper()

	if out.AccessToken != "access-token-xyz" {
		t.Fatalf("access token = %q", out.AccessToken)
	}
	if out.TrustLevel != "aal1" {
		t.Fatalf("trust level = %q, want aal1", out.TrustLevel)
	}
	if out.StepUpRequired {
		t.Fatal("step_up_required = true, want false")
	}
	if strings.TrimSpace(out.RefreshToken) == "" {
		t.Fatal("refresh token is empty")
	}
	if transactor.runCalls != 1 {
		t.Fatalf("transactor run calls = %d, want 1", transactor.runCalls)
	}
	if accountRepo.resolveCalls != 1 || accountRepo.upsertCalls != 1 {
		t.Fatalf("account repo calls = resolve:%d upsert:%d", accountRepo.resolveCalls, accountRepo.upsertCalls)
	}
	if accountRepo.lastUpsertAccID != accountRepo.accountID {
		t.Fatalf("upsert account id = %q", accountRepo.lastUpsertAccID)
	}
	if accountRepo.lastIdentity.Subject != "google-sub-123" {
		t.Fatalf("identity subject = %q", accountRepo.lastIdentity.Subject)
	}
	if roleAssigner.ensureCalls != 1 || roleAssigner.lastRoleKey != "base" {
		t.Fatalf("ensure role calls=%d role=%q", roleAssigner.ensureCalls, roleAssigner.lastRoleKey)
	}
	if roleAssigner.lastAccountID != accountRepo.accountID {
		t.Fatalf("ensure role account id = %q", roleAssigner.lastAccountID)
	}
	if sessionStore.createCalls != 1 || sessionStore.lastSession.AccountID != accountRepo.accountID {
		t.Fatalf("session create/account mismatch")
	}
	if sessionStore.lastSession.ID == "" || sessionStore.lastSession.RefreshTokenHash == "" {
		t.Fatal("session id or refresh token hash is empty")
	}
	if sessionStore.lastSession.RefreshTokenHash == out.RefreshToken {
		t.Fatal("refresh token hash must not equal plain refresh token")
	}
	if tokenIssuer.issueCalls != 1 {
		t.Fatalf("token issue calls = %d, want 1", tokenIssuer.issueCalls)
	}
	if tokenIssuer.lastReq.AccountID != accountRepo.accountID {
		t.Fatalf("token issuer account id = %q", tokenIssuer.lastReq.AccountID)
	}
	if tokenIssuer.lastReq.SessionID != sessionStore.lastSession.ID {
		t.Fatalf("token issuer session id = %q, want %q", tokenIssuer.lastReq.SessionID, sessionStore.lastSession.ID)
	}
	if tokenIssuer.lastReq.TrustLevel != "aal1" {
		t.Fatalf("token issuer trust level = %q", tokenIssuer.lastReq.TrustLevel)
	}
}
