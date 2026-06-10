package usecase

import (
	"time"

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
