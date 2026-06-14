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
// along with gopos-api. If not, see https://www.gnu.org/licenses/.

package usecase

import (
	"time"

	"pos-go/internal/modules/supplier/domain"
)

const (
	defaultListPage    = 1
	defaultListPerPage = 10
	maxListPerPage     = 50
	defaultLookupLimit = 20
	maxLookupLimit     = 50
)

type SupplierIDGenerator func() (domain.SupplierID, error)

type Clock func() time.Time

type SupplierContactInput struct {
	Phone   string
	Email   string
	Address string
	Notes   string
}

type SupplierResult struct {
	ID             string
	Name           string
	NormalizedName string
	Phone          string
	Email          string
	Address        string
	Notes          string
	IsActive       bool
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SupplierLookupResult struct {
	ID    string
	Name  string
	Phone string
	Email string
}

func mapSupplierResult(supplier domain.Supplier) SupplierResult {
	return SupplierResult{
		ID:             string(supplier.ID()),
		Name:           supplier.Name(),
		NormalizedName: string(supplier.NormalizedName()),
		Phone:          supplier.Phone(),
		Email:          supplier.Email(),
		Address:        supplier.Address(),
		Notes:          supplier.Notes(),
		IsActive:       supplier.IsActive(),
		Status:         string(supplier.Status()),
		CreatedAt:      supplier.CreatedAt(),
		UpdatedAt:      supplier.UpdatedAt(),
	}
}

func mapSupplierResults(suppliers []domain.Supplier) []SupplierResult {
	results := make([]SupplierResult, 0, len(suppliers))
	for _, supplier := range suppliers {
		results = append(results, mapSupplierResult(supplier))
	}

	return results
}

func mapSupplierLookupResults(suppliers []domain.Supplier) []SupplierLookupResult {
	results := make([]SupplierLookupResult, 0, len(suppliers))
	for _, supplier := range suppliers {
		results = append(results, SupplierLookupResult{
			ID:    string(supplier.ID()),
			Name:  supplier.Name(),
			Phone: supplier.Phone(),
			Email: supplier.Email(),
		})
	}

	return results
}

func ensureClock(clock Clock) Clock {
	if clock != nil {
		return clock
	}

	return time.Now
}

func mapContactInput(input SupplierContactInput) domain.SupplierContact {
	return domain.SupplierContact{
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		Notes:   input.Notes,
	}
}
