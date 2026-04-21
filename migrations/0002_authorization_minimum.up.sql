CREATE TABLE roles (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    key text NOT NULL,
    name text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT roles_key_unique UNIQUE (key)
);

CREATE TABLE permissions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    key text NOT NULL,
    name text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT permissions_key_unique UNIQUE (key)
);

CREATE TABLE account_roles (
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    role_id uuid NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT account_roles_account_id_role_id_unique UNIQUE (account_id, role_id)
);

CREATE TABLE role_permissions (
    role_id uuid NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id uuid NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT role_permissions_role_id_permission_id_unique UNIQUE (role_id, permission_id)
);

CREATE INDEX account_roles_account_id_idx
    ON account_roles (account_id);

CREATE INDEX account_roles_role_id_idx
    ON account_roles (role_id);

CREATE INDEX role_permissions_role_id_idx
    ON role_permissions (role_id);

CREATE INDEX role_permissions_permission_id_idx
    ON role_permissions (permission_id);
