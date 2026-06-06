package ports

import (
	"context"
	"errors"

	"pos-go/internal/modules/capability/domain"
)

var ErrCapabilityNotFound = errors.New("capability not found")

type CapabilityRepository interface {
	List(ctx context.Context) ([]domain.Capability, error)
	Get(ctx context.Context, key string) (domain.Capability, error)
	Save(ctx context.Context, capability domain.Capability) error
}
