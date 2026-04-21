package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

type fakeSessionStore struct {
	createCalls int
	lastSession domain.Session
}

func (f *fakeSessionStore) Create(ctx context.Context, session domain.Session) error {
	_ = ctx
	f.createCalls++
	f.lastSession = session
	return nil
}

type fakeTokenIssuer struct {
	issueCalls int
	lastReq    ports.AccessTokenRequest
	token      string
	exp        time.Time
}

func (f *fakeTokenIssuer) IssueAccessToken(ctx context.Context, req ports.AccessTokenRequest) (string, time.Time, error) {
	_ = ctx
	f.issueCalls++
	f.lastReq = req
	return f.token, f.exp, nil
}

type fakeTransactor struct {
	runCalls int
}

func (f *fakeTransactor) RunInTx(ctx context.Context, fn func(context.Context) error) error {
	f.runCalls++
	return fn(ctx)
}

type fakeGoogleCallbackRoleAssigner struct {
	ensureCalls   int
	lastAccountID string
	lastRoleKey   string
}

func (f *fakeGoogleCallbackRoleAssigner) EnsureRole(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.ensureCalls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return nil
}
