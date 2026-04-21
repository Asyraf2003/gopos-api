CREATE TABLE accounts (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE auth_identities (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    provider text NOT NULL,
    subject text NOT NULL,
    email text NOT NULL,
    email_verified boolean NOT NULL DEFAULT false,
    meta_json jsonb NOT NULL DEFAULT '{}'::jsonb,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT auth_identities_provider_subject_unique UNIQUE (provider, subject)
);

CREATE TABLE auth_sessions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    refresh_token_hash text NOT NULL,
    expires_at timestamptz NOT NULL,
    revoked_at timestamptz NULL,
    meta_json jsonb NOT NULL DEFAULT '{}'::jsonb,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX auth_identities_account_id_idx
    ON auth_identities (account_id);

CREATE INDEX auth_sessions_account_id_idx
    ON auth_sessions (account_id);

CREATE INDEX auth_sessions_expires_at_idx
    ON auth_sessions (expires_at);
