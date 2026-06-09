package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

type UpdateProduct struct {
	repository        ports.ProductRepository
	duplicateChecker  ports.ProductDuplicateChecker
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewUpdateProduct(
	repository ports.ProductRepository,
	duplicateChecker ports.ProductDuplicateChecker,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *UpdateProduct {
	return &UpdateProduct{
		repository:        repository,
		duplicateChecker:  duplicateChecker,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}

func (uc *UpdateProduct) Execute(
	ctx context.Context,
	cmd UpdateProductCommand,
) (UpdateProductResult, error) {
	product, err := uc.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return UpdateProductResult{}, err
	}

	_ = product

	return UpdateProductResult{}, nil
}
