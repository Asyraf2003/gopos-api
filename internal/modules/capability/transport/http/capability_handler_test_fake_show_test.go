package http

import (
	"context"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type fakeShowCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeShowCapabilityUsecase) Execute(ctx context.Context, key string) (domain.Capability, error) {
	_ = ctx
	f.fake.showCalls++
	if f.fake.err != nil {
		return domain.Capability{}, f.fake.err
	}

	capability, ok := f.fake.capabilities[key]
	if !ok {
		return domain.Capability{}, ports.ErrCapabilityNotFound
	}

	return capability, nil
}

var _ ShowCapabilityUsecase = fakeShowCapabilityUsecase{}
