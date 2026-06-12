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

type ProductVersionResponse struct {
	ProductID        string `json:"product_id"`
	RevisionNo       int    `json:"revision_no"`
	EventName        string `json:"event_name"`
	ChangedByActorID string `json:"changed_by_actor_id"`
	ChangeReason     string `json:"change_reason"`
	ChangedAt        string `json:"changed_at"`
}

func FromProductVersions(result productcatalogusecase.ListProductVersionsResult) []ProductVersionResponse {
	responses := make([]ProductVersionResponse, 0, len(result.Items))
	for _, item := range result.Items {
		responses = append(responses, ProductVersionResponse{
			ProductID:        item.ProductID,
			RevisionNo:       item.RevisionNo,
			EventName:        item.EventName,
			ChangedByActorID: item.ChangedByActorID,
			ChangeReason:     item.ChangeReason,
			ChangedAt:        formatRFC3339(item.ChangedAt),
		})
	}

	return responses
}
