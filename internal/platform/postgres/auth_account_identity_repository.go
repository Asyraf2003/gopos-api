package postgres

import (
	"context"
	"encoding/json"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountIdentityRepository struct {
	pool *pgxpool.Pool
}

func NewAccountIdentityRepository(pool *pgxpool.Pool) *AccountIdentityRepository {
	return &AccountIdentityRepository{pool: pool}
}

func (r *AccountIdentityRepository) ResolveOrCreateAccountByGoogle(ctx context.Context, claims ports.OIDCClaims) (string, error) {
	var accountID string

	err := r.queryRow(ctx,
		`SELECT account_id FROM auth_identities WHERE provider = $1 AND subject = $2`,
		claims.Provider, claims.Subject,
	).Scan(&accountID)
	if err == nil {
		return accountID, nil
	}
	if err != pgx.ErrNoRows {
		return "", err
	}

	err = r.queryRow(ctx,
		`SELECT id FROM accounts WHERE email = $1 LIMIT 1`,
		claims.Email,
	).Scan(&accountID)
	if err == nil {
		return accountID, nil
	}
	if err != pgx.ErrNoRows {
		return "", err
	}

	err = r.queryRow(ctx,
		`INSERT INTO accounts (email) VALUES ($1) RETURNING id`,
		claims.Email,
	).Scan(&accountID)
	if err != nil {
		return "", err
	}

	return accountID, nil
}

func (r *AccountIdentityRepository) UpsertIdentity(ctx context.Context, accountID string, identity domain.Identity) error {
	metaJSON, err := json.Marshal(identity.Meta)
	if err != nil {
		return err
	}

	return r.exec(ctx, `
		INSERT INTO auth_identities (
			account_id, provider, subject, email, email_verified, meta_json
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (provider, subject)
		DO UPDATE SET
			account_id = EXCLUDED.account_id,
			email = EXCLUDED.email,
			email_verified = EXCLUDED.email_verified,
			meta_json = EXCLUDED.meta_json,
			updated_at = now()
	`,
		accountID,
		string(identity.Provider),
		identity.Subject,
		identity.Email,
		identity.EmailVerified,
		metaJSON,
	)
}

func (r *AccountIdentityRepository) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *AccountIdentityRepository) exec(ctx context.Context, sql string, args ...any) error {
	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, args...)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}

var _ ports.AccountIdentityRepository = (*AccountIdentityRepository)(nil)
