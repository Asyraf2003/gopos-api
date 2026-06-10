package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type ListProducts struct {
	reader ports.ProductReader
}

func NewListProducts(reader ports.ProductReader) *ListProducts {
	return &ListProducts{
		reader: reader,
	}
}

func (uc *ListProducts) Execute(
	ctx context.Context,
	query ListProductsQuery,
) (ListProductsResult, error) {
	items, err := uc.reader.List(ctx, ports.ProductListQuery{
		Search:  query.Search,
		Status:  query.Status,
		Page:    query.Page,
		PerPage: query.PerPage,
	})
	if err != nil {
		return ListProductsResult{}, err
	}

	result := ListProductsResult{
		Items: make([]ListProductsItem, 0, len(items)),
	}
	for _, item := range items {
		result.Items = append(result.Items, ListProductsItem{
			ID:              item.ID,
			Code:            item.Code,
			Name:            item.Name,
			Brand:           item.Brand,
			Size:            item.Size,
			SalePriceRupiah: item.SalePriceRupiah,
			Status:          item.Status,
		})
	}

	return result, nil
}
