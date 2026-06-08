package servicecatalog

import (
	"time"

	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"
)

type ServiceCatalogItemResponse struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	NormalizedName      string   `json:"normalized_name"`
	DefaultPriceRupiah  int64    `json:"default_price_rupiah"`
	IsActive            bool     `json:"is_active"`
	Status              string   `json:"status"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
	AvailableOperations []string `json:"available_operations"`
}

func FromServiceCatalogItem(
	result servicecatalogusecase.ServiceCatalogItemResult,
) ServiceCatalogItemResponse {
	return ServiceCatalogItemResponse{
		ID:                  result.ID,
		Name:                result.Name,
		NormalizedName:      result.NormalizedName,
		DefaultPriceRupiah:  result.DefaultPriceRupiah,
		IsActive:            result.IsActive,
		Status:              result.Status,
		CreatedAt:           formatRFC3339(result.CreatedAt),
		UpdatedAt:           formatRFC3339(result.UpdatedAt),
		AvailableOperations: []string{},
	}
}

func FromServiceCatalogItems(
	results []servicecatalogusecase.ServiceCatalogItemResult,
) []ServiceCatalogItemResponse {
	responses := make([]ServiceCatalogItemResponse, 0, len(results))
	for _, result := range results {
		responses = append(responses, FromServiceCatalogItem(result))
	}

	return responses
}

func formatRFC3339(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format(time.RFC3339)
}
