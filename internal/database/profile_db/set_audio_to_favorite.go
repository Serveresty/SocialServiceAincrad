package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"context"
)

func SetAudioToFavorite(userID string, songID int) error {
	_, err := database.DB.Exec(context.Background(), `
	UPDATE users_songs SET songs_list = ARRAY_APPEND(users_songs.songs_list, CAST($1 AS bigint)) WHERE user_id = $2;
	`, songID, userID)
	//`UPDATE users_songs SET songs_list = ARRAY_APPEND(users_songs.songs_list, CAST(new_song_id AS bigint)) WHERE user_id = your_user_id;`
	if err != nil {
		return err
	}
	return nil
}
