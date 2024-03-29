package profiledb

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"context"
)

func GetAuthData(user *models.AuthUser) (int, error) {
	row := database.DB.QueryRow(context.Background(), `SELECT user_id, password FROM "users_data" WHERE email = $1`, user.Email)

	var id int
	var passwordHash string

	err := row.Scan(&id, &passwordHash)
	if err != nil {
		return 0, err
	}

	if ok := utils.CheckPasswordHash(user.Password, passwordHash); !ok {
		return 0, cerr.ErrHashError
	}

	return id, nil
}
