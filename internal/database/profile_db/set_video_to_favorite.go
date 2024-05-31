package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"context"
)

func SetVideoToFavorite(userID string, videoID int) error {
	_, err := database.DB.Exec(context.Background(), `
	UPDATE users_videos SET videos_list = ARRAY_APPEND(users_videos.videos_list, CAST($1 AS bigint)) WHERE user_id = $2;
	`, videoID, userID)
	if err != nil {
		return err
	}
	return nil
}
