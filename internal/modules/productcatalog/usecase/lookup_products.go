package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type LookupProducts struct {
	reader ports.ProductReader
}

func NewLookupProducts(reader ports.ProductReader) *LookupProducts {
	return &LookupProducts{
		reader: reader,
	}
}

func (uc *LookupProducts) Execute(
	ctx context.Context,
	query LookupProductsQuery,
) (LookupProductsResult, error) {
	_, err := uc.reader.Lookup(ctx, ports.ProductLookupQuery{
		Query:          query.Query,
		Limit:          query.Limit,
		IncludeDeleted: query.IncludeDeleted,
	})
	if err != nil {
		return LookupProductsResult{}, err
	}

	return LookupProductsResult{}, nil
}
