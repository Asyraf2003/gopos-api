package postgres

import (
	"context"
	"encoding/json"

	"pos-go/internal/modules/auth/domain"
)

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
