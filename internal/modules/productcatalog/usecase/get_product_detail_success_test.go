package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestGetProductDetailReturnsProductDetail(t *testing.T) {
	size := 12
	reorderPointQty := 8
	criticalThresholdQty := 3
	product, err := domain.NewProduct(domain.ProductInput{
		ID:                   "product-1",
		Code:                 "FLT-001",
		Name:                 "Filter Udara",
		Brand:                "Aspira",
		Size:                 &size,
		SalePriceRupiah:      30000,
		ReorderPointQty:      &reorderPointQty,
		CriticalThresholdQty: &criticalThresholdQty,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	usecase := NewGetProductDetail(&getProductDetailReaderDouble{
		found: product,
	})

	result, err := usecase.Execute(context.Background(), GetProductDetailQuery{
		ID: "product-1",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if result.ID != "product-1" {
		t.Fatalf("ID = %q, want product-1", result.ID)
	}
	if result.Code == nil || *result.Code != "FLT-001" {
		t.Fatalf("Code = %v, want FLT-001", result.Code)
	}
	if result.Name != "Filter Udara" {
		t.Fatalf("Name = %q, want Filter Udara", result.Name)
	}
	if result.NormalizedName != "filter udara" {
		t.Fatalf("NormalizedName = %q, want filter udara", result.NormalizedName)
	}
	if result.Brand != "Aspira" {
		t.Fatalf("Brand = %q, want Aspira", result.Brand)
	}
	if result.NormalizedBrand != "aspira" {
		t.Fatalf("NormalizedBrand = %q, want aspira", result.NormalizedBrand)
	}
	if result.Size == nil || *result.Size != 12 {
		t.Fatalf("Size = %v, want 12", result.Size)
	}
	if result.SalePriceRupiah != 30000 {
		t.Fatalf("SalePriceRupiah = %d, want 30000", result.SalePriceRupiah)
	}
	if result.ReorderPointQty == nil || *result.ReorderPointQty != 8 {
		t.Fatalf("ReorderPointQty = %v, want 8", result.ReorderPointQty)
	}
	if result.CriticalThresholdQty == nil || *result.CriticalThresholdQty != 3 {
		t.Fatalf("CriticalThresholdQty = %v, want 3", result.CriticalThresholdQty)
	}
	if result.Status != "active" {
		t.Fatalf("Status = %q, want active", result.Status)
	}
}