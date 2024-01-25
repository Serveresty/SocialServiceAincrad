package utilsdb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
	"fmt"
)

func IsUserRegistered(user *models.User) bool {
	query := `SELECT user_id FROM "users_data" WHERE username = $1 or email = $2 or phone = $3`
	rows, err := database.DB.Query(context.Background(), query, user.Username, user.Email, user.Phone)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer rows.Close()

	var arr []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return true
		}

		arr = append(arr, id)
	}

	if len(arr) > 0 {
		return true
	}

	return false
}
