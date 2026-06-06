package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type ShowCapability struct {
	repository ports.CapabilityRepository
}

func NewShowCapability(repository ports.CapabilityRepository) *ShowCapability {
	return &ShowCapability{repository: repository}
}

func (u *ShowCapability) Execute(ctx context.Context, key string) (domain.Capability, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return domain.Capability{}, errors.New("capability key is required")
	}

	return u.repository.Get(ctx, key)
}
