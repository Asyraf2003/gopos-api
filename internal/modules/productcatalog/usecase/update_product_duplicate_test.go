package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type fakeUpdateProductRepository struct {
	product *domain.Product
}

func (f *fakeUpdateProductRepository) Create(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeUpdateProductRepository) Update(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeUpdateProductRepository) FindByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return f.product, nil
}

type fakeUpdateProductDuplicateChecker struct {
	updateCalled bool
	productID    string
	candidate    ports.ProductDuplicateCandidate
}

func (f *fakeUpdateProductDuplicateChecker) CheckCreateDuplicate(
	_ context.Context,
	_ ports.ProductDuplicateCandidate,
) error {
	return nil
}

func (f *fakeUpdateProductDuplicateChecker) CheckUpdateDuplicate(
	_ context.Context,
	productID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	f.updateCalled = true
	f.productID = productID
	f.candidate = candidate

	return nil
}

func TestUpdateProductChecksDuplicateCandidate(t *testing.T) {
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_001",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	checker := &fakeUpdateProductDuplicateChecker{}
	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		checker,
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_001",
		Code:            "  PRD-002  ",
		Name:            "  Oli   Mesin ",
		Brand:           " Yamaha   Genuine ",
		Size:            domain.IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if !checker.updateCalled {
		t.Fatalf("update duplicate checker was not called")
	}
	if checker.productID != "prod_001" {
		t.Fatalf("productID = %q, want prod_001", checker.productID)
	}
	if checker.candidate.Code == nil || *checker.candidate.Code != "PRD-002" {
		t.Fatalf("candidate.Code = %v, want PRD-002", checker.candidate.Code)
	}
	if checker.candidate.NormalizedName != "oli mesin" {
		t.Fatalf("candidate.NormalizedName = %q, want oli mesin", checker.candidate.NormalizedName)
	}
	if checker.candidate.NormalizedBrand != "yamaha genuine" {
		t.Fatalf("candidate.NormalizedBrand = %q, want yamaha genuine", checker.candidate.NormalizedBrand)
	}
}
