package postgres

import "context"

func (r *PrincipalResolver) loadPermissions(ctx context.Context, accountID string) ([]string, error) {
	rows, err := r.query(ctx, `
		SELECT DISTINCT p.key
		FROM account_roles ar
		JOIN role_permissions rp ON rp.role_id = ar.role_id
		JOIN permissions p ON p.id = rp.permission_id
		WHERE ar.account_id = $1
		ORDER BY p.key
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var permissionKey string
		if err := rows.Scan(&permissionKey); err != nil {
			return nil, err
		}
		permissions = append(permissions, permissionKey)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
