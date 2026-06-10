package ports

import (
	"context"
	"errors"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

var (
	ErrProductNotFound          = errors.New("product not found")
	ErrDuplicateProductCode     = errors.New("duplicate product code")
	ErrDuplicateProductIdentity = errors.New("duplicate product identity")
)

type ProductListQuery struct {
	Search  string
	Status  string
	Page    int
	PerPage int
}

type ProductListItem struct {
	ID              string
	Code            *string
	Name            string
	Brand           string
	Size            *int
	SalePriceRupiah int64
	Status          string
}

type ProductLookupQuery struct{}

type ProductLookupItem struct{}

type ProductVersionRecord struct {
	ProductID        string
	RevisionNo       int
	EventName        string
	ChangedByActorID string
	ChangeReason     string
	ChangedAt        time.Time
}

type ProductAuditRecord struct {
	AggregateID string
	EventName   string
	Operation   string
	ActorID     string
	Reason      string
	OccurredAt  time.Time
	RevisionNo  int
}

type ProductAuditRecorder interface {
	RecordProductAudit(ctx context.Context, record ProductAuditRecord) error
}

type ProductDuplicateCandidate struct {
	Code            *string
	NormalizedName  string
	NormalizedBrand string
	Size            *int
}

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	FindByID(ctx context.Context, id string) (*domain.Product, error)
}

type ProductReader interface {
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	List(ctx context.Context, query ProductListQuery) ([]ProductListItem, error)
	Lookup(ctx context.Context, query ProductLookupQuery) ([]ProductLookupItem, error)
}

type ProductVersionRepository interface {
	Append(ctx context.Context, version ProductVersionRecord) error
	ListByProductID(ctx context.Context, productID string) ([]ProductVersionRecord, error)
}

type ProductDuplicateChecker interface {
	CheckCreateDuplicate(ctx context.Context, candidate ProductDuplicateCandidate) error
	CheckUpdateDuplicate(ctx context.Context, productID string, candidate ProductDuplicateCandidate) error
}
