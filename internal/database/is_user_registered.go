package database

import (
	"SocialServiceAincrad/models"
	"context"
)

func IsUserRegistered(user *models.User) bool {
	query := `INSERT INTO users_data (first_name, last_name, sex, username, email, phone) VALUES `
	args := user.First_name + user.Last_name + user.Sex + user.Username + user.Email + user.Phone
	_, err := DB.Exec(context.Background(), query, args)
	if err != nil {
		return true
	}

	return false
}
