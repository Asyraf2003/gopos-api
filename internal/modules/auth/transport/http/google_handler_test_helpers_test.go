package http

import (
	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type fakeGoogleFlow struct {
	startInput    authusecase.GoogleStartInput
	startOutput   authusecase.GoogleStartOutput
	startErr      error
	callbackInput authusecase.GoogleCallbackInput
	callbackOut   authusecase.GoogleCallbackOutput
	callbackErr   error
}

func (f *fakeGoogleFlow) GoogleStart(ctx echo.Context, in authusecase.GoogleStartInput) (authusecase.GoogleStartOutput, error) {
	_ = ctx
	f.startInput = in
	return f.startOutput, f.startErr
}

func (f *fakeGoogleFlow) GoogleCallback(ctx echo.Context, in authusecase.GoogleCallbackInput) (authusecase.GoogleCallbackOutput, error) {
	_ = ctx
	f.callbackInput = in
	return f.callbackOut, f.callbackErr
}
