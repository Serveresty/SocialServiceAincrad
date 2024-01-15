package database

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"context"
)

func GetAuthData(user *models.AuthUser) (string, error) {
	row := DB.QueryRow(context.Background(), `SELECT email, password FROM "users_data" WHERE email = $1`, user.Email)

	var email, passwordHash string

	err := row.Scan(&email, &passwordHash)
	if err != nil {
		return "", err
	}

	if ok := utils.CheckPasswordHash(user.Password, passwordHash); !ok {
		return "", cerr.HashError
	}

	return email, nil
}
