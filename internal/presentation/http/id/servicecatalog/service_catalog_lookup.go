package servicecatalog

import servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"

type ServiceCatalogLookupResponse struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	DefaultPriceRupiah int64  `json:"default_price_rupiah"`
}

func FromServiceCatalogLookup(
	result servicecatalogusecase.ServiceCatalogLookupResult,
) ServiceCatalogLookupResponse {
	return ServiceCatalogLookupResponse{
		ID:                 result.ID,
		Name:               result.Name,
		DefaultPriceRupiah: result.DefaultPriceRupiah,
	}
}

func FromServiceCatalogLookups(
	results []servicecatalogusecase.ServiceCatalogLookupResult,
) []ServiceCatalogLookupResponse {
	responses := make([]ServiceCatalogLookupResponse, 0, len(results))
	for _, result := range results {
		responses = append(responses, FromServiceCatalogLookup(result))
	}

	return responses
}
