package domain

import "time"

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Code() *string {
	if p.code == nil {
		return nil
	}
	code := *p.code
	return &code
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) NormalizedName() string {
	return p.normalizedName
}

func (p *Product) Brand() string {
	return p.brand
}

func (p *Product) NormalizedBrand() string {
	return p.normalizedBrand
}

func (p *Product) Size() *int {
	return copyIntPtr(p.size)
}

func (p *Product) SalePriceRupiah() int64 {
	return p.salePriceRupiah
}

func (p *Product) ReorderPointQty() *int {
	return copyIntPtr(p.reorderPointQty)
}

func (p *Product) CriticalThresholdQty() *int {
	return copyIntPtr(p.criticalThresholdQty)
}

func (p *Product) Status() ProductStatus {
	if p.deletedAt != nil {
		return ProductStatusDeleted
	}
	return ProductStatusActive
}

func (p *Product) DeletedAt() *time.Time {
	if p.deletedAt == nil {
		return nil
	}
	deletedAt := *p.deletedAt
	return &deletedAt
}
