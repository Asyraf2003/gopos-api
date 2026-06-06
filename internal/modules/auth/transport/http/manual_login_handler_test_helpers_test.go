package http

import (
	"context"

	authusecase "pos-go/internal/modules/auth/usecase"
)

type fakeManualLoginUsecase struct {
	lastInput authusecase.ManualLoginInput
	output    authusecase.ManualLoginOutput
	err       error
}

func (f *fakeManualLoginUsecase) Execute(
	ctx context.Context,
	in authusecase.ManualLoginInput,
) (authusecase.ManualLoginOutput, error) {
	_ = ctx
	f.lastInput = in
	return f.output, f.err
}
