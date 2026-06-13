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

package response

import (
	"encoding/json"
	"testing"
)

func TestSuccessIncludesEmptyMeta(t *testing.T) {
	body := Success(map[string]string{"id": "product-1"})

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal success envelope: %v", err)
	}

	var decoded struct {
		Success bool           `json:"success"`
		Data    map[string]any `json:"data"`
		Meta    map[string]any `json:"meta"`
	}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal success envelope: %v", err)
	}

	if !decoded.Success {
		t.Fatal("expected success=true")
	}
	if decoded.Meta == nil {
		t.Fatal("expected meta to be an empty object, got nil")
	}
	if len(decoded.Meta) != 0 {
		t.Fatalf("expected empty meta, got %#v", decoded.Meta)
	}
}

func TestSuccessWithMetaPreservesMeta(t *testing.T) {
	body := SuccessWithMeta("ok", map[string]any{"page": float64(1)})

	if body.Meta == nil {
		t.Fatal("expected meta")
	}

	meta, ok := body.Meta.(map[string]any)
	if !ok {
		t.Fatalf("expected map meta, got %T", body.Meta)
	}
	if meta["page"] != float64(1) {
		t.Fatalf("expected supplied meta, got %#v", meta)
	}
}

func TestSuccessWithMetaNormalizesNilMeta(t *testing.T) {
	body := SuccessWithMeta("ok", nil)

	meta, ok := body.Meta.(map[string]any)
	if !ok {
		t.Fatalf("expected map meta, got %T", body.Meta)
	}
	if len(meta) != 0 {
		t.Fatalf("expected empty meta, got %#v", meta)
	}
}
