package postgres

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceCatalogRepository struct {
	pool *pgxpool.Pool
}

func NewServiceCatalogRepository(pool *pgxpool.Pool) *ServiceCatalogRepository {
	return &ServiceCatalogRepository{pool: pool}
}

func (r *ServiceCatalogRepository) Create(ctx context.Context, item domain.ServiceCatalogItem) error {
	sql := `
		INSERT INTO service_catalog_items (
			id,
			name,
			normalized_name,
			default_price_rupiah,
			is_active,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.exec(ctx, sql, serviceCatalogItemArgs(item)...)
	return err
}

func (r *ServiceCatalogRepository) Update(ctx context.Context, item domain.ServiceCatalogItem) error {
	sql := `
		UPDATE service_catalog_items
		SET
			name = $2,
			normalized_name = $3,
			default_price_rupiah = $4,
			is_active = $5,
			updated_at = $7
		WHERE id = $1
	`

	_, err := r.exec(ctx, sql, serviceCatalogItemArgs(item)...)
	return err
}

func (r *ServiceCatalogRepository) FindByID(
	ctx context.Context,
	id domain.ServiceCatalogItemID,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE id = $1
	`, string(id))

	item, err := scanServiceCatalogItem(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ServiceCatalogItem{}, false, nil
	}
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	return item, true, nil
}

func (r *ServiceCatalogRepository) FindByNormalizedName(
	ctx context.Context,
	normalizedName domain.NormalizedName,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE normalized_name = $1
	`, string(normalizedName))

	item, err := scanServiceCatalogItem(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ServiceCatalogItem{}, false, nil
	}
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	return item, true, nil
}

func (r *ServiceCatalogRepository) List(
	ctx context.Context,
	filter ports.ListServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	args := []any{}
	conditions := []string{}
	nextArg := 1

	if strings.TrimSpace(filter.Query) != "" {
		conditions = append(conditions, "normalized_name LIKE $"+itoa(nextArg))
		args = append(args, "%"+domain.NormalizeName(filter.Query)+"%")
		nextArg++
	}

	switch filter.Status {
	case ports.ListStatusActive, "":
		conditions = append(conditions, "is_active = true")
	case ports.ListStatusInactive:
		conditions = append(conditions, "is_active = false")
	case ports.ListStatusAll:
	default:
		conditions = append(conditions, "is_active = true")
	}

	page := filter.Page
	if page <= 0 {
		page = 1
	}

	perPage := filter.PerPage
	if perPage <= 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage
	args = append(args, perPage, offset)
	limitArg := nextArg
	offsetArg := nextArg + 1

	sql := serviceCatalogItemSelectSQL()
	if len(conditions) > 0 {
		sql += "\n\t\tWHERE " + strings.Join(conditions, " AND ")
	}
	sql += "\n\t\tORDER BY normalized_name, id"
	sql += "\n\t\tLIMIT $" + itoa(limitArg) + " OFFSET $" + itoa(offsetArg)

	return r.findMany(ctx, sql, args...)
}

func (r *ServiceCatalogRepository) Lookup(
	ctx context.Context,
	filter ports.LookupServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	args := []any{}
	conditions := []string{}
	nextArg := 1

	if strings.TrimSpace(filter.Query) != "" {
		conditions = append(conditions, "normalized_name LIKE $"+itoa(nextArg))
		args = append(args, "%"+domain.NormalizeName(filter.Query)+"%")
		nextArg++
	}

	if filter.ActiveOnly {
		conditions = append(conditions, "is_active = true")
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit)
	limitArg := nextArg

	sql := serviceCatalogItemSelectSQL()
	if len(conditions) > 0 {
		sql += "\n\t\tWHERE " + strings.Join(conditions, " AND ")
	}
	sql += "\n\t\tORDER BY normalized_name, id"
	sql += "\n\t\tLIMIT $" + itoa(limitArg)

	return r.findMany(ctx, sql, args...)
}

func (r *ServiceCatalogRepository) SetActive(
	ctx context.Context,
	id domain.ServiceCatalogItemID,
	active bool,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE id = $1
		FOR UPDATE
	`, string(id))

	item, err := scanServiceCatalogItem(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ServiceCatalogItem{}, false, nil
	}
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	if active {
		item.Activate(nowUTC())
	} else {
		item.Deactivate(nowUTC())
	}

	err = r.Update(ctx, item)
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	return item, true, nil
}

func (r *ServiceCatalogRepository) findMany(
	ctx context.Context,
	sql string,
	args ...any,
) ([]domain.ServiceCatalogItem, error) {
	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []domain.ServiceCatalogItem{}
	for rows.Next() {
		item, err := scanServiceCatalogItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func itoa(value int) string {
	return strconvFormatInt(int64(value))
}

var _ ports.ServiceCatalogRepository = (*ServiceCatalogRepository)(nil)
