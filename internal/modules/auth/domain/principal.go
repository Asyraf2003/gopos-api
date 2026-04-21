package domain

type Principal struct {
	AccountID   string
	SessionID   string
	Roles       []string
	Permissions []string
	TrustLevel  string
}

func (p Principal) HasPermission(permissionKey string) bool {
	for _, permission := range p.Permissions {
		if permission == permissionKey {
			return true
		}
	}

	return false
}
