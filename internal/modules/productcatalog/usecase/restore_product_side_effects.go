package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

const productRestoredEventName = "product_restored"

func (uc *RestoreProduct) recordRestoreProductVersion(
	ctx context.Context,
	productID string,
	cmd RestoreProductCommand,
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
		EventName:        productRestoredEventName,
		ChangedByActorID: cmd.ActorID,
		ChangeReason:     cmd.Reason,
		ChangedAt:        occurredAt,
	}

	return revisionNo, uc.versionRepository.Append(ctx, version)
}

func (uc *RestoreProduct) recordRestoreProductAudit(
	ctx context.Context,
	productID string,
	cmd RestoreProductCommand,
	occurredAt time.Time,
	revisionNo int,
) error {
	return uc.auditRecorder.RecordProductAudit(ctx, ports.ProductAuditRecord{
		AggregateID: productID,
		EventName:   productRestoredEventName,
		Operation:   "restore",
		ActorID:     cmd.ActorID,
		Reason:      cmd.Reason,
		OccurredAt:  occurredAt,
		RevisionNo:  revisionNo,
	})
}
