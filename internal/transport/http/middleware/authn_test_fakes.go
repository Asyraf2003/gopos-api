package middleware

import (
	"context"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

type fakeAccessTokenVerifier struct {
	claims ports.AccessTokenClaims
	err    error
}

func (f *fakeAccessTokenVerifier) VerifyAccessToken(ctx context.Context, token string) (ports.AccessTokenClaims, error) {
	_ = ctx
	_ = token
	return f.claims, f.err
}

type fakePrincipalResolver struct {
	principal domain.Principal
	err       error
	lastInput ports.ResolvePrincipalInput
}

func (f *fakePrincipalResolver) Resolve(ctx context.Context, in ports.ResolvePrincipalInput) (domain.Principal, error) {
	_ = ctx
	f.lastInput = in
	return f.principal, f.err
}

type fakeSessionStatusChecker struct {
	active        bool
	err           error
	lastSessionID string
}

func (f *fakeSessionStatusChecker) IsSessionActive(ctx context.Context, sessionID string) (bool, error) {
	_ = ctx
	f.lastSessionID = sessionID
	return f.active, f.err
}
