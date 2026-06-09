package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestCreateProductSuccessPersistsProductAndChecksDuplicateCandidate(t *testing.T) {
	fixedNow := time.Date(2026, 6, 9, 12, 0, 0, 0, time.UTC)
	repository := &fakeProductRepository{}
	duplicateChecker := &fakeProductDuplicateChecker{}

	uc := NewCreateProduct(
		repository,
		duplicateChecker,
		fakeProductIDGenerator{id: "prod_001"},
		func() time.Time { return fixedNow },
	)

	result, err := uc.Execute(context.Background(), CreateProductCommand{
		Code:            "  PRD-001  ",
		Name:            "  Oli   Mesin ",
		Brand:           " Yamaha   Genuine ",
		Size:            domain.IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.created == nil {
		t.Fatalf("repository.Create was not called")
	}

	if !duplicateChecker.createCalled {
		t.Fatalf("duplicate checker was not called")
	}

	if duplicateChecker.candidate.Code == nil || *duplicateChecker.candidate.Code != "PRD-001" {
		t.Fatalf("candidate.Code = %v, want PRD-001", duplicateChecker.candidate.Code)
	}
	if duplicateChecker.candidate.NormalizedName != "oli mesin" {
		t.Fatalf("candidate.NormalizedName = %q, want oli mesin", duplicateChecker.candidate.NormalizedName)
	}
	if duplicateChecker.candidate.NormalizedBrand != "yamaha genuine" {
		t.Fatalf("candidate.NormalizedBrand = %q, want yamaha genuine", duplicateChecker.candidate.NormalizedBrand)
	}
	if duplicateChecker.candidate.Size == nil || *duplicateChecker.candidate.Size != 1000 {
		t.Fatalf("candidate.Size = %v, want 1000", duplicateChecker.candidate.Size)
	}

	if result.ID != "prod_001" {
		t.Fatalf("result.ID = %q, want prod_001", result.ID)
	}
	if result.Name != "Oli Mesin" {
		t.Fatalf("result.Name = %q, want Oli Mesin", result.Name)
	}
	if result.Brand != "Yamaha Genuine" {
		t.Fatalf("result.Brand = %q, want Yamaha Genuine", result.Brand)
	}
	if result.Status != string(domain.ProductStatusActive) {
		t.Fatalf("result.Status = %q, want active", result.Status)
	}
	if !result.CreatedAt.Equal(fixedNow) || !result.UpdatedAt.Equal(fixedNow) {
		t.Fatalf("result timestamps = %v/%v, want %v", result.CreatedAt, result.UpdatedAt, fixedNow)
	}
}

func TestCreateProductReturnsDuplicateCheckerError(t *testing.T) {
	duplicateErr := errors.New("duplicate failure")

	uc := NewCreateProduct(
		&fakeProductRepository{},
		&fakeProductDuplicateChecker{err: duplicateErr},
		fakeProductIDGenerator{id: "prod_002"},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), CreateProductCommand{
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})
	if !errors.Is(err, duplicateErr) {
		t.Fatalf("Execute() error = %v, want %v", err, duplicateErr)
	}
}
