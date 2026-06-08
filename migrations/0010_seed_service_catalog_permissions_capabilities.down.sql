DELETE FROM api_capabilities
WHERE key IN (
    'service_catalog.list',
    'service_catalog.create',
    'service_catalog.lookup',
    'service_catalog.show',
    'service_catalog.update',
    'service_catalog.activate',
    'service_catalog.deactivate'
);

DELETE FROM role_permissions
WHERE permission_id IN (
    SELECT id
    FROM permissions
    WHERE key IN ('service_catalog.read', 'service_catalog.manage')
)
AND role_id IN (
    SELECT id
    FROM roles
    WHERE key IN ('admin', 'cashier')
);

DELETE FROM permissions
WHERE key IN ('service_catalog.read', 'service_catalog.manage');
