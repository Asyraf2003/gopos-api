package postgres

import (
	"errors"

	"pos-go/internal/modules/productcatalog/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

var errProductRepositoryNotImplemented = errors.New("product postgres repository behavior not implemented")

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}

var _ ports.ProductRepository = (*ProductRepository)(nil)
var _ ports.ProductReader = (*ProductRepository)(nil)
var _ ports.ProductVersionRepository = (*ProductRepository)(nil)
var _ ports.ProductDuplicateChecker = (*ProductRepository)(nil)
