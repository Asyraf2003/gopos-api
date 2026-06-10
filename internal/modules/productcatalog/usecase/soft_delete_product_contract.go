package usecase

import "time"

type SoftDeleteProductCommand struct {
	ID      string
	ActorID string
	Reason  string
}

type SoftDeleteProductResult struct {
	ID         string
	Status     string
	DeletedAt  time.Time
	RevisionNo int
}
