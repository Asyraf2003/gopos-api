package postgres

import (
	"context"
	"time"

	"pos-go/internal/modules/system/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthChecker struct {
	pool *pgxpool.Pool
}

func NewHealthChecker(pool *pgxpool.Pool) *HealthChecker {
	return &HealthChecker{pool: pool}
}

func (h *HealthChecker) Check(ctx context.Context) error {
	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return h.pool.Ping(pingCtx)
}

var _ ports.HealthChecker = (*HealthChecker)(nil)
