DELETE FROM api_capabilities
WHERE key IN (
    'profile.self.show',
    'authz.profile.self.show',
    'auth.session.logout',
    'account.role.assign',
    'account.role.remove'
);
