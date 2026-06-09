package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
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

	candidate, err := domain.NewProduct(domain.ProductInput{
		ID:                   product.ID(),
		Code:                 cmd.Code,
		Name:                 cmd.Name,
		Brand:                cmd.Brand,
		Size:                 cmd.Size,
		SalePriceRupiah:      cmd.SalePriceRupiah,
		ReorderPointQty:      cmd.ReorderPointQty,
		CriticalThresholdQty: cmd.CriticalThresholdQty,
	})
	if err != nil {
		return UpdateProductResult{}, err
	}

	if err := uc.duplicateChecker.CheckUpdateDuplicate(ctx, product.ID(), ports.ProductDuplicateCandidate{
		Code:            candidate.Code(),
		NormalizedName:  candidate.NormalizedName(),
		NormalizedBrand: candidate.NormalizedBrand(),
		Size:            candidate.Size(),
	}); err != nil {
		return UpdateProductResult{}, err
	}

	return UpdateProductResult{}, nil
}
