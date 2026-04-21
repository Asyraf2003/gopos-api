package http

import (
	"context"

	authusecase "pos-go/internal/modules/auth/usecase"
)

type fakeAssignAccountRoleUsecase struct {
	lastAccountID string
	lastRoleKey   string
	calls         int
	err           error
}

func (f *fakeAssignAccountRoleUsecase) Execute(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.calls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return f.err
}

type fakeRemoveAccountRoleUsecase struct {
	lastAccountID string
	lastRoleKey   string
	calls         int
	err           error
}

func (f *fakeRemoveAccountRoleUsecase) Execute(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.calls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return f.err
}

func newAccountRoleHandlerForTest(
	assignErr error,
	removeErr error,
) (*AccountRoleHandler, *fakeAssignAccountRoleUsecase, *fakeRemoveAccountRoleUsecase) {
	assignUsecase := &fakeAssignAccountRoleUsecase{err: assignErr}
	removeUsecase := &fakeRemoveAccountRoleUsecase{err: removeErr}

	handler := NewAccountRoleHandler(assignUsecase, removeUsecase)

	if removeErr == authusecase.ErrBaseRoleRemovalNotAllowed {
		return handler, assignUsecase, removeUsecase
	}

	return handler, assignUsecase, removeUsecase
}
