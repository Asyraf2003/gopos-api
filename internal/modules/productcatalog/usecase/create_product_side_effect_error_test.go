package usecase

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCreateProductReturnsVersionRepositoryError(t *testing.T) {
	versionErr := errors.New("version append failure")
	repository := &fakeProductRepository{}

	uc := NewCreateProduct(
		repository,
		&fakeProductDuplicateChecker{},
		&fakeProductVersionRepository{err: versionErr},
		&fakeProductAuditRecorder{},
		fakeProductIDGenerator{id: "prod_003"},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), CreateProductCommand{
		Name:            "Filter Oli",
		Brand:           "Yamaha",
		SalePriceRupiah: 45000,
	})
	if !errors.Is(err, versionErr) {
		t.Fatalf("Execute() error = %v, want %v", err, versionErr)
	}
	if repository.created == nil {
		t.Fatalf("repository.Create was not called")
	}
}

func TestCreateProductReturnsAuditRecorderError(t *testing.T) {
	auditErr := errors.New("audit record failure")
	auditRecorder := &fakeProductAuditRecorder{err: auditErr}

	uc := NewCreateProduct(
		&fakeProductRepository{},
		&fakeProductDuplicateChecker{},
		&fakeProductVersionRepository{},
		auditRecorder,
		fakeProductIDGenerator{id: "prod_004"},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), CreateProductCommand{
		Name:            "Kampas Rem",
		Brand:           "Honda",
		SalePriceRupiah: 65000,
	})
	if !errors.Is(err, auditErr) {
		t.Fatalf("Execute() error = %v, want %v", err, auditErr)
	}
	if len(auditRecorder.records) != 1 {
		t.Fatalf("audit record count = %d, want 1", len(auditRecorder.records))
	}
}
