package database

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"context"
)

func GetAuthData(user *models.AuthUser) (int, string, error) {
	row := DB.QueryRow(context.Background(), `SELECT id, username, password FROM "users_data" WHERE email = $1`, user.Email)

	var id int
	var passwordHash string
	var username string

	err := row.Scan(&id, &username, &passwordHash)
	if err != nil {
		return 0, "", err
	}

	if ok := utils.CheckPasswordHash(user.Password, passwordHash); !ok {
		return 0, "", cerr.HashError
	}

	return id, username, nil
}
