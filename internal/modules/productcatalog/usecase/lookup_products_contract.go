package usecase

type LookupProductsQuery struct {
	Query          string
	Limit          int
	IncludeDeleted bool
}

type LookupProductsResult struct {
	Items []LookupProductsItem
}

type LookupProductsItem struct {
	ID              string
	Code            *string
	Name            string
	Brand           string
	Size            *int
	SalePriceRupiah int64
	Status          string
}
