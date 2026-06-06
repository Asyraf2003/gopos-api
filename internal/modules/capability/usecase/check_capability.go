package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type CheckCapability struct {
	repository ports.CapabilityRepository
}

func NewCheckCapability(repository ports.CapabilityRepository) *CheckCapability {
	return &CheckCapability{repository: repository}
}

func (u *CheckCapability) Execute(ctx context.Context, key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("capability key is required")
	}

	capability, err := u.repository.Get(ctx, key)
	if err != nil {
		return err
	}

	if !capability.Enabled {
		return domain.ErrCapabilityDisabled
	}

	return nil
}
