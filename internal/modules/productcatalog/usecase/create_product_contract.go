package usecase

import "time"

type CreateProductCommand struct {
	Code                 string
	Name                 string
	Brand                string
	Size                 *int
	SalePriceRupiah      int64
	ReorderPointQty      *int
	CriticalThresholdQty *int
	ActorID              string
	Reason               string
}

type CreateProductResult struct {
	ID                   string
	Code                 *string
	Name                 string
	NormalizedName       string
	Brand                string
	NormalizedBrand      string
	Size                 *int
	SalePriceRupiah      int64
	ReorderPointQty      *int
	CriticalThresholdQty *int
	Status               string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
