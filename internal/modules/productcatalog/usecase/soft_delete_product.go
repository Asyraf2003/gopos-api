package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type SoftDeleteProduct struct {
	repository        ports.ProductRepository
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewSoftDeleteProduct(
	repository ports.ProductRepository,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *SoftDeleteProduct {
	return &SoftDeleteProduct{
		repository:        repository,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}

func (uc *SoftDeleteProduct) Execute(
	ctx context.Context,
	cmd SoftDeleteProductCommand,
) (SoftDeleteProductResult, error) {
	product, err := uc.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return SoftDeleteProductResult{}, err
	}

	deletedAt := uc.now()
	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: cmd.ActorID,
		Reason:           cmd.Reason,
	}); err != nil {
		return SoftDeleteProductResult{}, err
	}

	return SoftDeleteProductResult{
		ID:        product.ID(),
		Status:    string(product.Status()),
		DeletedAt: deletedAt,
	}, nil
}
