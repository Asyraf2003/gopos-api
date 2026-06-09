package usecase

import (
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func createProductResultFromDomain(product *domain.Product, now time.Time) CreateProductResult {
	return CreateProductResult{
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
		CreatedAt:            now,
		UpdatedAt:            now,
	}
}
