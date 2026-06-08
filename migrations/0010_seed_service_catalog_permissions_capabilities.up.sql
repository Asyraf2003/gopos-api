INSERT INTO permissions (key, name)
VALUES
    ('service_catalog.read', 'Read service catalog'),
    ('service_catalog.manage', 'Manage service catalog')
ON CONFLICT (key) DO UPDATE SET
    name = EXCLUDED.name;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key = 'service_catalog.read'
WHERE r.key IN ('admin', 'cashier')
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key = 'service_catalog.manage'
WHERE r.key = 'admin'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO api_capabilities (
    key,
    domain,
    operation,
    method,
    path,
    default_enabled,
    enabled,
    required_permission,
    risk_level,
    audit_required,
    idempotency_required,
    owner_package,
    test_proof,
    disabled_reason
)
VALUES
    (
        'service_catalog.list',
        'service_catalog',
        'list',
        'GET',
        '/api/service-catalog/items',
        true,
        true,
        'service_catalog.read',
        'low',
        false,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'service_catalog.create',
        'service_catalog',
        'create',
        'POST',
        '/api/service-catalog/items',
        true,
        true,
        'service_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'service_catalog.lookup',
        'service_catalog',
        'lookup',
        'GET',
        '/api/service-catalog/items/lookup',
        true,
        true,
        'service_catalog.read',
        'low',
        false,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'service_catalog.show',
        'service_catalog',
        'show',
        'GET',
        '/api/service-catalog/items/:id',
        true,
        true,
        'service_catalog.read',
        'low',
        false,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'service_catalog.update',
        'service_catalog',
        'update',
        'PUT',
        '/api/service-catalog/items/:id',
        true,
        true,
        'service_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'service_catalog.activate',
        'service_catalog',
        'activate',
        'POST',
        '/api/service-catalog/items/:id/activate',
        true,
        true,
        'service_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'service_catalog.deactivate',
        'service_catalog',
        'deactivate',
        'POST',
        '/api/service-catalog/items/:id/deactivate',
        true,
        true,
        'service_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/servicecatalog/transport/http',
        'internal/modules/servicecatalog/transport/http plus route capability audit proof',
        NULL
    )
ON CONFLICT (key) DO UPDATE SET
    domain = EXCLUDED.domain,
    operation = EXCLUDED.operation,
    method = EXCLUDED.method,
    path = EXCLUDED.path,
    default_enabled = EXCLUDED.default_enabled,
    enabled = EXCLUDED.enabled,
    required_permission = EXCLUDED.required_permission,
    risk_level = EXCLUDED.risk_level,
    audit_required = EXCLUDED.audit_required,
    idempotency_required = EXCLUDED.idempotency_required,
    owner_package = EXCLUDED.owner_package,
    test_proof = EXCLUDED.test_proof,
    disabled_reason = EXCLUDED.disabled_reason,
    updated_at = now();
