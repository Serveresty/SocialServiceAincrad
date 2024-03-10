package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func GetAudiosListById(id int) ([]models.Audio, error) {
	rows, err := database.DB.Query(context.Background(), `SELECT s.song_id, s.name, s.author FROM users_songs u JOIN songs s ON u.songs_list @> ARRAY[s.song_id] WHERE u.user_id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Audio
	for rows.Next() {
		var song models.Audio
		err := rows.Scan(&song.Id, &song.Name, &song.Author)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
