package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type ListProductVersions struct {
	versions ports.ProductVersionRepository
}

func NewListProductVersions(versions ports.ProductVersionRepository) *ListProductVersions {
	return &ListProductVersions{
		versions: versions,
	}
}

func (uc *ListProductVersions) Execute(
	ctx context.Context,
	query ListProductVersionsQuery,
) (ListProductVersionsResult, error) {
	records, err := uc.versions.ListByProductID(ctx, query.ProductID)
	if err != nil {
		return ListProductVersionsResult{}, err
	}

	result := ListProductVersionsResult{
		Items: make([]ListProductVersionItem, 0, len(records)),
	}
	for _, record := range records {
		result.Items = append(result.Items, ListProductVersionItem{
			ProductID:        record.ProductID,
			RevisionNo:       record.RevisionNo,
			EventName:        record.EventName,
			ChangedByActorID: record.ChangedByActorID,
			ChangeReason:     record.ChangeReason,
			ChangedAt:        record.ChangedAt,
		})
	}

	return result, nil
}
