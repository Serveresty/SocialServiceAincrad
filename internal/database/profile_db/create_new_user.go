package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
	"strconv"
)

func CreateUser(user *models.User) error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO "users_data" (first_name, last_name, sex, username, email, phone, password) VALUES($1,$2,$3,$4,$5,$6,$7)`,
		user.First_name, user.Last_name, user.Sex, user.Email, user.Email, user.Phone, user.Password)
	if err != nil {
		return err
	}

	row := database.DB.QueryRow(context.Background(), `SELECT user_id FROM "users_data" WHERE email = $1`, user.Email)

	err = row.Scan(&user.Id)
	if err != nil {
		return err
	}

	id := strconv.Itoa(user.Id)
	user.Username = id

	_, err = database.DB.Exec(context.Background(), `UPDATE "users_data" SET username=$1 WHERE user_id = $1`, user.Id)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(context.Background(), `INSERT INTO "users_roles" (user_id, role_id) VALUES($1, $2)`, user.Id, 1)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(context.Background(), `INSERT INTO "users_songs" (user_id) VALUES($1)`, user.Id)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(context.Background(), `INSERT INTO "users_videos" (user_id) VALUES($1)`, user.Id)
	if err != nil {
		return err
	}

	err = generateSettings(user.Id)
	if err != nil {
		return err
	}

	return nil
}

func generateSettings(id int) error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO "social" (user_id) VALUES($1)`, id)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(context.Background(), `INSERT INTO "info" (user_id) VALUES($1)`, id)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(context.Background(), `INSERT INTO "general_settings" (user_id, front_style, censure_filter, language) VALUES($1, $2, $3, $4)`, id, 1, false, "Русский")
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(context.Background(), `INSERT INTO "privacy_settings" (user_id, saved_photos, groups, audio, video, friends, posts, comments, messages) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`, id, 1, 1, 1, 1, 1, 1, 1, 1)
	if err != nil {
		return err
	}
	return nil
}
