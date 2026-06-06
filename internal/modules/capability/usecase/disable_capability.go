package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/capability/ports"
)

type DisableCapability struct {
	repository ports.CapabilityRepository
}

func NewDisableCapability(repository ports.CapabilityRepository) *DisableCapability {
	return &DisableCapability{repository: repository}
}

func (u *DisableCapability) Execute(ctx context.Context, key string, reason string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("capability key is required")
	}

	capability, err := u.repository.Get(ctx, key)
	if err != nil {
		return err
	}

	return u.repository.Save(ctx, capability.Disable(reason))
}
