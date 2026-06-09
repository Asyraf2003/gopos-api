package domain

import "strings"

func NewProduct(input ProductInput) (*Product, error) {
	id := strings.TrimSpace(input.ID)
	if id == "" {
		return nil, ErrProductIDRequired
	}

	name := normalizeDisplayText(input.Name)
	if name == "" {
		return nil, ErrProductNameRequired
	}

	brand := normalizeDisplayText(input.Brand)
	if brand == "" {
		return nil, ErrProductBrandRequired
	}

	if input.SalePriceRupiah <= 0 {
		return nil, ErrProductSalePriceMustBePositive
	}

	if err := validateThreshold(input.ReorderPointQty, input.CriticalThresholdQty); err != nil {
		return nil, err
	}

	return &Product{
		id:                   id,
		code:                 normalizeCode(input.Code),
		name:                 name,
		normalizedName:       normalizeSearchText(name),
		brand:                brand,
		normalizedBrand:      normalizeSearchText(brand),
		size:                 copyIntPtr(input.Size),
		salePriceRupiah:      input.SalePriceRupiah,
		reorderPointQty:      copyIntPtr(input.ReorderPointQty),
		criticalThresholdQty: copyIntPtr(input.CriticalThresholdQty),
	}, nil
}
