package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestListProductsMapsReaderItems(t *testing.T) {
	code := "BRG-001"
	size := 250
	reader := &listProductsReaderDouble{
		listItems: []ports.ProductListItem{
			{
				ID:              "product-1",
				Code:            &code,
				Name:            "Kopi Arabika",
				Brand:           "Acme",
				Size:            &size,
				SalePriceRupiah: 15000,
				Status:          "active",
			},
		},
	}
	usecase := NewListProducts(reader)

	result, err := usecase.Execute(context.Background(), ListProductsQuery{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 1 {
		t.Fatalf("result Items length = %d, want %d", len(result.Items), 1)
	}

	item := result.Items[0]
	if item.ID != "product-1" {
		t.Fatalf("item ID = %q, want %q", item.ID, "product-1")
	}
	if item.Code == nil || *item.Code != "BRG-001" {
		t.Fatalf("item Code = %v, want %q", item.Code, "BRG-001")
	}
	if item.Name != "Kopi Arabika" {
		t.Fatalf("item Name = %q, want %q", item.Name, "Kopi Arabika")
	}
	if item.Brand != "Acme" {
		t.Fatalf("item Brand = %q, want %q", item.Brand, "Acme")
	}
	if item.Size == nil || *item.Size != 250 {
		t.Fatalf("item Size = %v, want %d", item.Size, 250)
	}
	if item.SalePriceRupiah != 15000 {
		t.Fatalf("item SalePriceRupiah = %d, want %d", item.SalePriceRupiah, 15000)
	}
	if item.Status != "active" {
		t.Fatalf("item Status = %q, want %q", item.Status, "active")
	}
}
