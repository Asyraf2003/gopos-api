package domain

import (
	"testing"
	"time"
)

func TestNewProductTrimsAndNormalizesNameBrandAndCode(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_001",
		Code:            "  BRG-001  ",
		Name:            "  Oli   Mesin  ",
		Brand:           "  Yamaha   Genuine ",
		Size:            IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	if product.Code() == nil || *product.Code() != "BRG-001" {
		t.Fatalf("Code() = %v, want BRG-001", product.Code())
	}
	if product.Name() != "Oli Mesin" {
		t.Fatalf("Name() = %q, want %q", product.Name(), "Oli Mesin")
	}
	if product.NormalizedName() != "oli mesin" {
		t.Fatalf("NormalizedName() = %q, want %q", product.NormalizedName(), "oli mesin")
	}
	if product.Brand() != "Yamaha Genuine" {
		t.Fatalf("Brand() = %q, want %q", product.Brand(), "Yamaha Genuine")
	}
	if product.NormalizedBrand() != "yamaha genuine" {
		t.Fatalf("NormalizedBrand() = %q, want %q", product.NormalizedBrand(), "yamaha genuine")
	}
}

func TestNewProductConvertsBlankCodeToNil(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_002",
		Code:            "   ",
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	if product.Code() != nil {
		t.Fatalf("Code() = %v, want nil", product.Code())
	}
}

func TestNewProductRejectsInvalidNameBrandAndPrice(t *testing.T) {
	tests := []struct {
		name  string
		input ProductInput
	}{
		{
			name: "blank name",
			input: ProductInput{
				ID:              "prod_blank_name",
				Name:            " ",
				Brand:           "NGK",
				SalePriceRupiah: 25000,
			},
		},
		{
			name: "blank brand",
			input: ProductInput{
				ID:              "prod_blank_brand",
				Name:            "Busi",
				Brand:           " ",
				SalePriceRupiah: 25000,
			},
		},
		{
			name: "zero sale price",
			input: ProductInput{
				ID:              "prod_zero_price",
				Name:            "Busi",
				Brand:           "NGK",
				SalePriceRupiah: 0,
			},
		},
		{
			name: "negative sale price",
			input: ProductInput{
				ID:              "prod_negative_price",
				Name:            "Busi",
				Brand:           "NGK",
				SalePriceRupiah: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewProduct(tt.input); err == nil {
				t.Fatalf("NewProduct() error = nil, want error")
			}
		})
	}
}

func TestNewProductValidatesThresholdPair(t *testing.T) {
	tests := []struct {
		name  string
		input ProductInput
	}{
		{
			name: "reorder without critical",
			input: ProductInput{
				ID:              "prod_reorder_only",
				Name:            "Kampas Rem",
				Brand:           "Honda",
				SalePriceRupiah: 40000,
				ReorderPointQty: IntPtr(10),
			},
		},
		{
			name: "critical without reorder",
			input: ProductInput{
				ID:                   "prod_critical_only",
				Name:                 "Kampas Rem",
				Brand:                "Honda",
				SalePriceRupiah:      40000,
				CriticalThresholdQty: IntPtr(3),
			},
		},
		{
			name: "negative reorder",
			input: ProductInput{
				ID:                   "prod_negative_reorder",
				Name:                 "Kampas Rem",
				Brand:                "Honda",
				SalePriceRupiah:      40000,
				ReorderPointQty:      IntPtr(-1),
				CriticalThresholdQty: IntPtr(0),
			},
		},
		{
			name: "critical greater than reorder",
			input: ProductInput{
				ID:                   "prod_bad_threshold_order",
				Name:                 "Kampas Rem",
				Brand:                "Honda",
				SalePriceRupiah:      40000,
				ReorderPointQty:      IntPtr(5),
				CriticalThresholdQty: IntPtr(6),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewProduct(tt.input); err == nil {
				t.Fatalf("NewProduct() error = nil, want error")
			}
		})
	}
}

func TestProductSoftDeleteAndRestoreLifecycle(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_003",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 9, 10, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: "actor_admin",
		Reason:           "duplicate test product",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}

	if product.Status() != ProductStatusDeleted {
		t.Fatalf("Status() = %v, want %v", product.Status(), ProductStatusDeleted)
	}
	if product.DeletedAt() == nil || !product.DeletedAt().Equal(deletedAt) {
		t.Fatalf("DeletedAt() = %v, want %v", product.DeletedAt(), deletedAt)
	}

	if err := product.Restore(); err != nil {
		t.Fatalf("Restore() error = %v", err)
	}

	if product.Status() != ProductStatusActive {
		t.Fatalf("Status() = %v, want %v", product.Status(), ProductStatusActive)
	}
	if product.DeletedAt() != nil {
		t.Fatalf("DeletedAt() = %v, want nil", product.DeletedAt())
	}
}
