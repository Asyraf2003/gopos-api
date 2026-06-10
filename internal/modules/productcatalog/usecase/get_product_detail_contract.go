package usecase

type GetProductDetailQuery struct {
	ID string
}

type GetProductDetailResult struct {
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
}
