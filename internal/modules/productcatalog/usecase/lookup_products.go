package usecase

import "pos-go/internal/modules/productcatalog/ports"

type LookupProducts struct {
	reader ports.ProductReader
}

func NewLookupProducts(reader ports.ProductReader) *LookupProducts {
	return &LookupProducts{
		reader: reader,
	}
}
