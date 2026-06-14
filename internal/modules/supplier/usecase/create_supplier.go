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
	"context"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"
)

type CreateSupplier struct {
	repo  ports.SupplierRepository
	newID SupplierIDGenerator
	now   Clock
}

type CreateSupplierCommand struct {
	Name    string
	Contact SupplierContactInput
}

func NewCreateSupplier(repo ports.SupplierRepository, newID SupplierIDGenerator, now Clock) CreateSupplier {
	return CreateSupplier{
		repo:  repo,
		newID: newID,
		now:   ensureClock(now),
	}
}

func (uc CreateSupplier) Execute(ctx context.Context, cmd CreateSupplierCommand) (SupplierResult, error) {
	if uc.newID == nil {
		return SupplierResult{}, ErrMissingSupplierIDGenerator
	}

	normalizedName := domain.NormalizeName(cmd.Name)
	existing, found, err := uc.repo.FindByNormalizedName(ctx, normalizedName)
	if err != nil {
		return SupplierResult{}, err
	}
	if found && existing.IsActive() {
		return SupplierResult{}, ErrDuplicateSupplierActiveName
	}

	id, err := uc.newID()
	if err != nil {
		return SupplierResult{}, err
	}

	supplier, err := domain.NewSupplier(id, cmd.Name, mapContactInput(cmd.Contact), uc.now())
	if err != nil {
		return SupplierResult{}, err
	}

	if err := uc.repo.Create(ctx, supplier); err != nil {
		return SupplierResult{}, err
	}

	return mapSupplierResult(supplier), nil
}
