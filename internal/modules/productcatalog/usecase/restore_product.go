package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

type RestoreProduct struct {
	repository        ports.ProductRepository
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewRestoreProduct(
	repository ports.ProductRepository,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *RestoreProduct {
	return &RestoreProduct{
		repository:        repository,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}

func (uc *RestoreProduct) Execute(
	ctx context.Context,
	cmd RestoreProductCommand,
) (RestoreProductResult, error) {
	product, err := uc.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return RestoreProductResult{}, err
	}

	if err := product.Restore(); err != nil {
		return RestoreProductResult{}, err
	}

	if err := uc.repository.Update(ctx, product); err != nil {
		return RestoreProductResult{}, err
	}

	restoredAt := uc.now()
	return RestoreProductResult{
		ID:         product.ID(),
		Status:     string(product.Status()),
		RestoredAt: restoredAt,
	}, nil
}
