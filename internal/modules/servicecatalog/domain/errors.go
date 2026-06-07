package domain

import "errors"

var (
	ErrInvalidServiceCatalogItemID             = errors.New("service catalog item id is required")
	ErrBlankServiceCatalogItemName             = errors.New("service catalog item name is required")
	ErrInvalidServiceCatalogItemNormalizedName = errors.New("service catalog item normalized name is required")
	ErrInvalidServiceCatalogItemDefaultPrice   = errors.New("service catalog item default price must be greater than zero")
)
