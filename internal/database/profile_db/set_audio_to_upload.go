package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func SetAudioToUpload(song models.Audio, hash string, filename string) (int, error) {
	var id int
	err := database.DB.QueryRow(context.Background(), `INSERT INTO "songs" (name, author, filename, hash) VALUES($1, $2, $3, $4) RETURNING song_id`, song.Name, song.Author, filename, hash).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
