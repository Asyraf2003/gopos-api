package usecase

import "pos-go/internal/modules/productcatalog/ports"

type ListProducts struct {
	reader ports.ProductReader
}

func NewListProducts(reader ports.ProductReader) *ListProducts {
	return &ListProducts{
		reader: reader,
	}
}
