package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrProductIDRequired              = errors.New("product id is required")
	ErrProductNameRequired            = errors.New("product name is required")
	ErrProductBrandRequired           = errors.New("product brand is required")
	ErrProductSalePriceMustBePositive = errors.New("product sale price must be greater than zero")
	ErrProductThresholdPairRequired   = errors.New("product reorder point and critical threshold must be both null or both filled")
	ErrProductThresholdNegative       = errors.New("product threshold must be non-negative")
	ErrProductCriticalAboveReorder    = errors.New("product critical threshold must not exceed reorder point")
	ErrProductDeleteTimeRequired      = errors.New("product delete time is required")
	ErrProductAlreadyDeleted          = errors.New("product is already deleted")
	ErrProductNotDeleted              = errors.New("product is not deleted")
)

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

func NewProduct(input ProductInput) (*Product, error) {
	id := strings.TrimSpace(input.ID)
	if id == "" {
		return nil, ErrProductIDRequired
	}

	name := normalizeDisplayText(input.Name)
	if name == "" {
		return nil, ErrProductNameRequired
	}

	brand := normalizeDisplayText(input.Brand)
	if brand == "" {
		return nil, ErrProductBrandRequired
	}

	if input.SalePriceRupiah <= 0 {
		return nil, ErrProductSalePriceMustBePositive
	}

	if err := validateThreshold(input.ReorderPointQty, input.CriticalThresholdQty); err != nil {
		return nil, err
	}

	return &Product{
		id:                   id,
		code:                 normalizeCode(input.Code),
		name:                 name,
		normalizedName:       normalizeSearchText(name),
		brand:                brand,
		normalizedBrand:      normalizeSearchText(brand),
		size:                 copyIntPtr(input.Size),
		salePriceRupiah:      input.SalePriceRupiah,
		reorderPointQty:      copyIntPtr(input.ReorderPointQty),
		criticalThresholdQty: copyIntPtr(input.CriticalThresholdQty),
	}, nil
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Code() *string {
	if p.code == nil {
		return nil
	}

	code := *p.code
	return &code
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) NormalizedName() string {
	return p.normalizedName
}

func (p *Product) Brand() string {
	return p.brand
}

func (p *Product) NormalizedBrand() string {
	return p.normalizedBrand
}

func (p *Product) Size() *int {
	return copyIntPtr(p.size)
}

func (p *Product) SalePriceRupiah() int64 {
	return p.salePriceRupiah
}

func (p *Product) ReorderPointQty() *int {
	return copyIntPtr(p.reorderPointQty)
}

func (p *Product) CriticalThresholdQty() *int {
	return copyIntPtr(p.criticalThresholdQty)
}

func (p *Product) Status() ProductStatus {
	if p.deletedAt != nil {
		return ProductStatusDeleted
	}

	return ProductStatusActive
}

func (p *Product) DeletedAt() *time.Time {
	if p.deletedAt == nil {
		return nil
	}

	deletedAt := *p.deletedAt
	return &deletedAt
}

func (p *Product) SoftDelete(input DeleteInput) error {
	if p.deletedAt != nil {
		return ErrProductAlreadyDeleted
	}

	if input.DeletedAt.IsZero() {
		return ErrProductDeleteTimeRequired
	}

	deletedAt := input.DeletedAt
	p.deletedAt = &deletedAt
	p.deletedByActorID = strings.TrimSpace(input.DeletedByActorID)
	p.deleteReason = strings.TrimSpace(input.Reason)

	return nil
}

func (p *Product) Restore() error {
	if p.deletedAt == nil {
		return ErrProductNotDeleted
	}

	p.deletedAt = nil
	p.deletedByActorID = ""
	p.deleteReason = ""

	return nil
}

func normalizeCode(code string) *string {
	trimmed := strings.TrimSpace(code)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}

func normalizeDisplayText(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func normalizeSearchText(value string) string {
	return strings.ToLower(normalizeDisplayText(value))
}

func validateThreshold(reorderPointQty, criticalThresholdQty *int) error {
	if (reorderPointQty == nil) != (criticalThresholdQty == nil) {
		return ErrProductThresholdPairRequired
	}

	if reorderPointQty == nil {
		return nil
	}

	if *reorderPointQty < 0 || *criticalThresholdQty < 0 {
		return ErrProductThresholdNegative
	}

	if *criticalThresholdQty > *reorderPointQty {
		return ErrProductCriticalAboveReorder
	}

	return nil
}

func copyIntPtr(value *int) *int {
	if value == nil {
		return nil
	}

	copied := *value
	return &copied
}

func IntPtr(value int) *int {
	return &value
}
