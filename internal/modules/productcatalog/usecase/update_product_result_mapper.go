package usecase

import (
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func updateProductResultFromDomain(
	product *domain.Product,
	updatedAt time.Time,
	revisionNo int,
) UpdateProductResult {
	return UpdateProductResult{
		ID:                   product.ID(),
		Code:                 product.Code(),
		Name:                 product.Name(),
		NormalizedName:       product.NormalizedName(),
		Brand:                product.Brand(),
		NormalizedBrand:      product.NormalizedBrand(),
		Size:                 product.Size(),
		SalePriceRupiah:      product.SalePriceRupiah(),
		ReorderPointQty:      product.ReorderPointQty(),
		CriticalThresholdQty: product.CriticalThresholdQty(),
		Status:               string(product.Status()),
		UpdatedAt:            updatedAt,
		RevisionNo:           revisionNo,
	}
}
