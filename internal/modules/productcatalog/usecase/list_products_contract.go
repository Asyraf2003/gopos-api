package usecase

type ListProductsQuery struct {
	Search  string
	Status  string
	Page    int
	PerPage int
}

type ListProductsResult struct {
	Items []ListProductsItem
}

type ListProductsItem struct {
	ID              string
	Code            *string
	Name            string
	Brand           string
	Size            *int
	SalePriceRupiah int64
	Status          string
}
