package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type fakeUpdateProductVersionRepository struct {
	existing []ports.ProductVersionRecord
	appended []ports.ProductVersionRecord
	err      error
}

func (f *fakeUpdateProductVersionRepository) Append(
	_ context.Context,
	version ports.ProductVersionRecord,
) error {
	f.appended = append(f.appended, version)

	return f.err
}

func (f *fakeUpdateProductVersionRepository) ListByProductID(
	_ context.Context,
	_ string,
) ([]ports.ProductVersionRecord, error) {
	return f.existing, nil
}
