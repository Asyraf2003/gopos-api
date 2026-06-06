package usecase

import (
	"context"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type ListCapabilities struct {
	repository ports.CapabilityRepository
}

func NewListCapabilities(repository ports.CapabilityRepository) *ListCapabilities {
	return &ListCapabilities{repository: repository}
}

func (u *ListCapabilities) Execute(ctx context.Context) ([]domain.Capability, error) {
	return u.repository.List(ctx)
}
