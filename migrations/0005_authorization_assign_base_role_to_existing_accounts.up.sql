INSERT INTO account_roles (account_id, role_id)
SELECT a.id, r.id
FROM accounts a
CROSS JOIN roles r
WHERE r.key = 'base'
ON CONFLICT (account_id, role_id) DO NOTHING;
