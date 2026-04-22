package postgres

import "context"

func (r *PrincipalResolver) loadRoles(ctx context.Context, accountID string) ([]string, error) {
	rows, err := r.query(ctx, `
		SELECT r.key
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1
		ORDER BY r.key
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleKey string
		if err := rows.Scan(&roleKey); err != nil {
			return nil, err
		}
		roles = append(roles, roleKey)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
