package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

const productDeletedEventName = "product_deleted"

func (uc *SoftDeleteProduct) recordSoftDeleteProductVersion(
	ctx context.Context,
	productID string,
	cmd SoftDeleteProductCommand,
	occurredAt time.Time,
) (int, error) {
	versions, err := uc.versionRepository.ListByProductID(ctx, productID)
	if err != nil {
		return 0, err
	}

	revisionNo := len(versions) + 1
	version := ports.ProductVersionRecord{
		ProductID:        productID,
		RevisionNo:       revisionNo,
		EventName:        productDeletedEventName,
		ChangedByActorID: cmd.ActorID,
		ChangeReason:     cmd.Reason,
		ChangedAt:        occurredAt,
	}

	return revisionNo, uc.versionRepository.Append(ctx, version)
}

func (uc *SoftDeleteProduct) recordSoftDeleteProductAudit(
	ctx context.Context,
	productID string,
	cmd SoftDeleteProductCommand,
	occurredAt time.Time,
	revisionNo int,
) error {
	return uc.auditRecorder.RecordProductAudit(ctx, ports.ProductAuditRecord{
		AggregateID: productID,
		EventName:   productDeletedEventName,
		Operation:   "delete",
		ActorID:     cmd.ActorID,
		Reason:      cmd.Reason,
		OccurredAt:  occurredAt,
		RevisionNo:  revisionNo,
	})
}
