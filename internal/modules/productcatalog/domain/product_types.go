package domain

import "time"

type ProductStatus string

const (
	ProductStatusActive  ProductStatus = "active"
	ProductStatusDeleted ProductStatus = "deleted"
)

type ProductInput struct {
	ID                   string
	Code                 string
	Name                 string
	Brand                string
	Size                 *int
	SalePriceRupiah      int64
	ReorderPointQty      *int
	CriticalThresholdQty *int
}

type DeleteInput struct {
	DeletedAt        time.Time
	DeletedByActorID string
	Reason           string
}

type Product struct {
	id                   string
	code                 *string
	name                 string
	normalizedName       string
	brand                string
	normalizedBrand      string
	size                 *int
	salePriceRupiah      int64
	reorderPointQty      *int
	criticalThresholdQty *int
	deletedAt            *time.Time
	deletedByActorID     string
	deleteReason         string
}

func IntPtr(value int) *int {
	return &value
}
