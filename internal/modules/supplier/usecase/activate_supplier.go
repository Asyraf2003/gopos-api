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
	"strings"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"
)

type ActivateSupplier struct {
	repo ports.SupplierRepository
}

type ActivateSupplierCommand struct {
	ID string
}

func NewActivateSupplier(repo ports.SupplierRepository) ActivateSupplier {
	return ActivateSupplier{repo: repo}
}

func (uc ActivateSupplier) Execute(ctx context.Context, cmd ActivateSupplierCommand) (SupplierResult, error) {
	id := domain.SupplierID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateSupplierID(id); err != nil {
		return SupplierResult{}, err
	}

	supplier, found, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return SupplierResult{}, err
	}
	if !found {
		return SupplierResult{}, ErrSupplierNotFound
	}

	existing, duplicateFound, err := uc.repo.FindByNormalizedName(ctx, supplier.NormalizedName())
	if err != nil {
		return SupplierResult{}, err
	}
	if duplicateFound && existing.ID() != supplier.ID() && existing.IsActive() {
		return SupplierResult{}, ErrDuplicateSupplierActiveName
	}

	activated, found, err := uc.repo.SetActive(ctx, id, true)
	if err != nil {
		return SupplierResult{}, err
	}
	if !found {
		return SupplierResult{}, ErrSupplierNotFound
	}

	return mapSupplierResult(activated), nil
}
