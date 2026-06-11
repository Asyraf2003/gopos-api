// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

//go:build integration

package postgres

import (
	"testing"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"

	"github.com/google/uuid"
)

func newDuplicateCheckerProduct(t *testing.T, code string, name string) *domain.Product {
	t.Helper()

	input := domain.ProductInput{
		ID:              uuid.NewString(),
		Name:            name,
		Brand:           "Honda",
		SalePriceRupiah: 40000,
	}
	if code != "" {
		input.Code = code
	}

	product, err := domain.NewProduct(input)
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	return product
}

func newDuplicateCheckerCandidate(
	t *testing.T,
	code string,
	name string,
) ports.ProductDuplicateCandidate {
	t.Helper()

	product := newDuplicateCheckerProduct(t, code, name)
	return duplicateCandidateFromProduct(product)
}

func duplicateCandidateFromProduct(product *domain.Product) ports.ProductDuplicateCandidate {
	return ports.ProductDuplicateCandidate{
		Code:            product.Code(),
		NormalizedName:  product.NormalizedName(),
		NormalizedBrand: product.NormalizedBrand(),
		Size:            product.Size(),
	}
}
