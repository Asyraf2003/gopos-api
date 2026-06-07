package domain

import (
	"strings"
	"time"
)

func RestoreServiceCatalogItem(
	id ServiceCatalogItemID,
	name string,
	normalizedName NormalizedName,
	defaultPriceRupiah MoneyRupiah,
	isActive bool,
	createdAt time.Time,
	updatedAt time.Time,
) (ServiceCatalogItem, error) {
	id = ServiceCatalogItemID(strings.TrimSpace(string(id)))
	normalizedName = NormalizedName(strings.TrimSpace(string(normalizedName)))

	if err := ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItem{}, err
	}

	if err := ValidateServiceCatalogItemName(name); err != nil {
		return ServiceCatalogItem{}, err
	}

	if normalizedName == "" {
		return ServiceCatalogItem{}, ErrInvalidServiceCatalogItemNormalizedName
	}

	if err := ValidateServiceCatalogItemDefaultPrice(defaultPriceRupiah); err != nil {
		return ServiceCatalogItem{}, err
	}

	return ServiceCatalogItem{
		id:                 id,
		name:               normalizeDisplayName(name),
		normalizedName:     normalizedName,
		defaultPriceRupiah: defaultPriceRupiah,
		isActive:           isActive,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
	}, nil
}
