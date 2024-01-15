package database

import (
	"SocialServiceAincrad/models"
	"context"
)

func CreateUser(user *models.User) error {
	_, err := DB.Exec(context.Background(), `INSERT INTO "users_data" (first_name, last_name, sex, username, email, phone, password) VALUES($1,$2,$3,$4,$5,$6,$7)`,
		user.First_name, user.Last_name, user.Sex, user.Username, user.Email, user.Phone, user.Password)
	if err != nil {
		return err
	}

	row := DB.QueryRow(context.Background(), `SELECT user_id FROM "users_data" WHERE email = $1`, user.Email)

	err = row.Scan(&user.Id)
	if err != nil {
		return err
	}

	_, err = DB.Exec(context.Background(), `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`, user.Id, 1)
	if err != nil {
		return err
	}

	return nil
}
