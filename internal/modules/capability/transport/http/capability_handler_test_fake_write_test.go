package http

import (
	"context"

	"pos-go/internal/modules/capability/ports"
)

type fakeEnableCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeEnableCapabilityUsecase) Execute(ctx context.Context, key string) error {
	_ = ctx
	f.fake.enableCalls++
	capability, ok := f.fake.capabilities[key]
	if !ok {
		return ports.ErrCapabilityNotFound
	}
	f.fake.capabilities[key] = capability.Enable()

	return f.fake.err
}

type fakeDisableCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeDisableCapabilityUsecase) Execute(ctx context.Context, key string, reason string) error {
	_ = ctx
	f.fake.disableCalls++
	f.fake.lastDisableReason = reason
	capability, ok := f.fake.capabilities[key]
	if !ok {
		return ports.ErrCapabilityNotFound
	}
	f.fake.capabilities[key] = capability.Disable(reason)

	return f.fake.err
}

var _ EnableCapabilityUsecase = fakeEnableCapabilityUsecase{}
var _ DisableCapabilityUsecase = fakeDisableCapabilityUsecase{}
