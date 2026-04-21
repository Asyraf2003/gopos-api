INSERT INTO roles (key, name)
VALUES
    ('base', 'Base'),
    ('cashier', 'Cashier'),
    ('admin', 'Admin')
ON CONFLICT (key) DO NOTHING;

INSERT INTO permissions (key, name)
VALUES
    ('auth.session.refresh', 'Refresh auth session'),
    ('auth.session.logout', 'Logout auth session'),
    ('profile.self.read', 'Read own profile'),
    ('sale.order.create', 'Create sale order'),
    ('sale.order.read', 'Read sale order'),
    ('payment.create', 'Create payment'),
    ('inventory.manage', 'Manage inventory'),
    ('report.read', 'Read reports'),
    ('account.role.assign', 'Assign account roles')
ON CONFLICT (key) DO NOTHING;
