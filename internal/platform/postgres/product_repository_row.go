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

package postgres

type productScanner interface {
	Scan(dest ...any) error
}

func productSelectSQL() string {
	return `
		SELECT
			id,
			kode_barang,
			nama_barang,
			nama_barang_normalized,
			merek,
			merek_normalized,
			ukuran,
			harga_jual,
			reorder_point_qty,
			critical_threshold_qty,
			deleted_at,
			deleted_by_actor_id,
			delete_reason
		FROM products
	`
}
