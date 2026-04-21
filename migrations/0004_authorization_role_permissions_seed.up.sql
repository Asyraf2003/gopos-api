INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key IN (
    'auth.session.refresh',
    'auth.session.logout',
    'profile.self.read'
)
WHERE r.key = 'base'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key IN (
    'auth.session.refresh',
    'auth.session.logout',
    'profile.self.read',
    'sale.order.create',
    'sale.order.read',
    'payment.create'
)
WHERE r.key = 'cashier'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key IN (
    'auth.session.refresh',
    'auth.session.logout',
    'profile.self.read',
    'sale.order.create',
    'sale.order.read',
    'payment.create',
    'inventory.manage',
    'report.read',
    'account.role.assign'
)
WHERE r.key = 'admin'
ON CONFLICT (role_id, permission_id) DO NOTHING;
