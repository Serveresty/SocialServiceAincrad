package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"context"
)

func GetUserRoles(id int) ([]string, error) {
	rows, err := database.DB.Query(context.Background(), `SELECT role_name FROM roles JOIN users_roles ON roles.role_id=users_roles.role_id WHERE users_roles.user_id=$1`, id)
	if err != nil {
		return []string{}, err
	}

	var roles []string
	for rows.Next() {

		var role string

		err = rows.Scan(&role)
		if err != nil {
			return []string{}, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
