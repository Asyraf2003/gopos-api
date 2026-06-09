package domain

import "errors"

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
