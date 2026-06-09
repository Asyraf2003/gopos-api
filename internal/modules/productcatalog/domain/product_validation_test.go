package domain

import "testing"

func TestNewProductRejectsInvalidNameBrandAndPrice(t *testing.T) {
	tests := []struct {
		name  string
		input ProductInput
	}{
		{name: "blank name", input: ProductInput{ID: "prod_blank_name", Name: " ", Brand: "NGK", SalePriceRupiah: 25000}},
		{name: "blank brand", input: ProductInput{ID: "prod_blank_brand", Name: "Busi", Brand: " ", SalePriceRupiah: 25000}},
		{name: "zero sale price", input: ProductInput{ID: "prod_zero_price", Name: "Busi", Brand: "NGK", SalePriceRupiah: 0}},
		{name: "negative sale price", input: ProductInput{ID: "prod_negative_price", Name: "Busi", Brand: "NGK", SalePriceRupiah: -1}},
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
			input: ProductInput{ID: "prod_reorder_only", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, ReorderPointQty: IntPtr(10)},
		},
		{
			name: "critical without reorder",
			input: ProductInput{ID: "prod_critical_only", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, CriticalThresholdQty: IntPtr(3)},
		},
		{
			name: "negative reorder",
			input: ProductInput{ID: "prod_negative_reorder", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, ReorderPointQty: IntPtr(-1), CriticalThresholdQty: IntPtr(0)},
		},
		{
			name: "critical greater than reorder",
			input: ProductInput{ID: "prod_bad_threshold_order", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, ReorderPointQty: IntPtr(5), CriticalThresholdQty: IntPtr(6)},
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
