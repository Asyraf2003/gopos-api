package usecase

import "time"

type UpdateProductCommand struct {
	ID                   string
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

type UpdateProductResult struct {
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
	UpdatedAt            time.Time
	RevisionNo           int
}
