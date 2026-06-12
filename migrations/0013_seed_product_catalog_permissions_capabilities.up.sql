-- Copyright (C) 2026 Asyraf Mubarak
--
-- This file is part of gopos-api.
--
-- gopos-api is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, version 3 only.
--
-- gopos-api is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
-- GNU Affero General Public License for more details.
--
-- You should have received a copy of the GNU Affero General Public License
-- along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

INSERT INTO permissions (key, name)
VALUES
    ('product_catalog.read', 'Read product catalog'),
    ('product_catalog.manage', 'Manage product catalog')
ON CONFLICT (key) DO UPDATE SET
    name = EXCLUDED.name;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key = 'product_catalog.read'
WHERE r.key IN ('admin', 'cashier')
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key = 'product_catalog.manage'
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
        'product_catalog.list',
        'product_catalog',
        'list',
        'GET',
        '/api/products',
        true,
        true,
        'product_catalog.read',
        'low',
        false,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.create',
        'product_catalog',
        'create',
        'POST',
        '/api/products',
        true,
        true,
        'product_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.lookup',
        'product_catalog',
        'lookup',
        'GET',
        '/api/products/lookup',
        true,
        true,
        'product_catalog.read',
        'low',
        false,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.show',
        'product_catalog',
        'show',
        'GET',
        '/api/products/:id',
        true,
        true,
        'product_catalog.read',
        'low',
        false,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.update',
        'product_catalog',
        'update',
        'PUT',
        '/api/products/:id',
        true,
        true,
        'product_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.delete',
        'product_catalog',
        'delete',
        'DELETE',
        '/api/products/:id',
        true,
        true,
        'product_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.restore',
        'product_catalog',
        'restore',
        'PATCH',
        '/api/products/:id/restore',
        true,
        true,
        'product_catalog.manage',
        'medium',
        true,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
        NULL
    ),
    (
        'product_catalog.versions',
        'product_catalog',
        'versions',
        'GET',
        '/api/products/:id/versions',
        true,
        true,
        'product_catalog.read',
        'low',
        false,
        false,
        'internal/modules/productcatalog/transport/http',
        'internal/modules/productcatalog/transport/http plus route capability audit proof',
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
