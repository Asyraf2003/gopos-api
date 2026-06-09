package domain

import "strings"

func (p *Product) SoftDelete(input DeleteInput) error {
	if p.deletedAt != nil {
		return ErrProductAlreadyDeleted
	}

	if input.DeletedAt.IsZero() {
		return ErrProductDeleteTimeRequired
	}

	deletedAt := input.DeletedAt
	p.deletedAt = &deletedAt
	p.deletedByActorID = strings.TrimSpace(input.DeletedByActorID)
	p.deleteReason = strings.TrimSpace(input.Reason)

	return nil
}

func (p *Product) Restore() error {
	if p.deletedAt == nil {
		return ErrProductNotDeleted
	}

	p.deletedAt = nil
	p.deletedByActorID = ""
	p.deleteReason = ""

	return nil
}
