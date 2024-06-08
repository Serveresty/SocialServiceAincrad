package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"context"
)

func AddFriend(adderID string, claimerID string) error {
	status := "wait"
	_, err := database.DB.Exec(context.Background(), `INSERT INTO "friends" (first, second, status_id) VALUES($1, $2, (SELECT status_id FROM friend_status WHERE status_name = $3))`, adderID, claimerID, status)
	if err != nil {
		return err
	}

	return nil
}
