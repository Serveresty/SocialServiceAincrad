package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func GetAudioById(id string) (models.Audio, error) {
	row := database.DB.QueryRow(context.Background(), `SELECT song_id, name, author, filename, hash FROM "songs" WHERE song_id = $1`, id)
	var audio models.Audio

	err := row.Scan(&audio.Id, &audio.Name, &audio.Author, &audio.Filename, &audio.Hash)
	if err != nil {
		return models.Audio{}, err
	}

	return audio, nil
}
