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

package productcatalog

import productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"

type ProductLookupResponse struct {
	ID              string  `json:"id"`
	Code            *string `json:"kode_barang"`
	Name            string  `json:"nama_barang"`
	Brand           string  `json:"merek"`
	Size            *int    `json:"ukuran"`
	SalePriceRupiah int64   `json:"harga_jual"`
	Status          string  `json:"status"`
}

func FromProductLookup(result productcatalogusecase.LookupProductsResult) []ProductLookupResponse {
	responses := make([]ProductLookupResponse, 0, len(result.Items))
	for _, item := range result.Items {
		responses = append(responses, ProductLookupResponse{
			ID:              item.ID,
			Code:            item.Code,
			Name:            item.Name,
			Brand:           item.Brand,
			Size:            item.Size,
			SalePriceRupiah: item.SalePriceRupiah,
			Status:          item.Status,
		})
	}

	return responses
}
