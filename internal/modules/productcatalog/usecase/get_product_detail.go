package usecase

import "pos-go/internal/modules/productcatalog/ports"

type GetProductDetail struct {
	reader ports.ProductReader
}

func NewGetProductDetail(reader ports.ProductReader) *GetProductDetail {
	return &GetProductDetail{
		reader: reader,
	}
}
