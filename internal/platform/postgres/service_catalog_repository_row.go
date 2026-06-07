package postgres

import (
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
)

type serviceCatalogItemScanner interface {
	Scan(dest ...any) error
}

func serviceCatalogItemSelectSQL() string {
	return `
		SELECT
			id,
			name,
			normalized_name,
			default_price_rupiah,
			is_active,
			created_at,
			updated_at
		FROM service_catalog_items
	`
}

func scanServiceCatalogItem(row serviceCatalogItemScanner) (domain.ServiceCatalogItem, error) {
	var id string
	var name string
	var normalizedName string
	var defaultPriceRupiah int64
	var isActive bool
	var createdAt time.Time
	var updatedAt time.Time

	err := row.Scan(
		&id,
		&name,
		&normalizedName,
		&defaultPriceRupiah,
		&isActive,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return domain.ServiceCatalogItem{}, err
	}

	return domain.RestoreServiceCatalogItem(
		domain.ServiceCatalogItemID(id),
		name,
		domain.NormalizedName(normalizedName),
		domain.MoneyRupiah(defaultPriceRupiah),
		isActive,
		createdAt,
		updatedAt,
	)
}

func serviceCatalogItemArgs(item domain.ServiceCatalogItem) []any {
	return []any{
		string(item.ID()),
		item.Name(),
		string(item.NormalizedName()),
		int64(item.DefaultPriceRupiah()),
		item.IsActive(),
		item.CreatedAt(),
		item.UpdatedAt(),
	}
}
