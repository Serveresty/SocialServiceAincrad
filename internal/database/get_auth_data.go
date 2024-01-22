package database

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"context"
)

func GetAuthData(user *models.AuthUser) (int, error) {
	row := DB.QueryRow(context.Background(), `SELECT id, password FROM "users_data" WHERE email = $1`, user.Email)

	var id int
	var passwordHash string

	err := row.Scan(&id, &passwordHash)
	if err != nil {
		return 0, err
	}

	if ok := utils.CheckPasswordHash(user.Password, passwordHash); !ok {
		return 0, cerr.HashError
	}

	return id, nil
}
