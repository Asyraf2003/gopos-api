package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/capability/ports"
)

type EnableCapability struct {
	repository ports.CapabilityRepository
}

func NewEnableCapability(repository ports.CapabilityRepository) *EnableCapability {
	return &EnableCapability{repository: repository}
}

func (u *EnableCapability) Execute(ctx context.Context, key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("capability key is required")
	}

	capability, err := u.repository.Get(ctx, key)
	if err != nil {
		return err
	}

	return u.repository.Save(ctx, capability.Enable())
}
