CREATE TABLE api_capabilities (
    key text PRIMARY KEY,
    domain text NOT NULL,
    operation text NOT NULL,
    method text NOT NULL,
    path text NOT NULL,
    default_enabled boolean NOT NULL,
    enabled boolean NOT NULL,
    required_permission text NOT NULL,
    risk_level text NOT NULL,
    audit_required boolean NOT NULL,
    idempotency_required boolean NOT NULL,
    owner_package text NOT NULL,
    test_proof text NOT NULL,
    disabled_reason text NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT api_capabilities_risk_level_check
        CHECK (risk_level IN ('low', 'medium', 'high')),
    CONSTRAINT api_capabilities_method_check
        CHECK (method IN ('GET', 'POST', 'PUT', 'PATCH', 'DELETE'))
);

CREATE INDEX api_capabilities_domain_operation_idx
    ON api_capabilities (domain, operation);

CREATE INDEX api_capabilities_required_permission_idx
    ON api_capabilities (required_permission);
