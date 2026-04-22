package http

import (
	"context"

	authusecase "pos-go/internal/modules/auth/usecase"
)

type fakeRefreshTokenUsecase struct {
	lastInput authusecase.RefreshTokenInput
	output    authusecase.RefreshTokenOutput
	err       error
}

func (f *fakeRefreshTokenUsecase) Execute(ctx context.Context, in authusecase.RefreshTokenInput) (authusecase.RefreshTokenOutput, error) {
	_ = ctx
	f.lastInput = in
	return f.output, f.err
}
