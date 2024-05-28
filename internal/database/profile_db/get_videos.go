package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func GetVideosListById(id int) ([]models.Video, error) {
	rows, err := database.DB.Query(context.Background(), `SELECT s.video_id, s.title, s.description, s.created_at FROM users_videos u JOIN videos s ON u.videos_list @> ARRAY[s.video_id] WHERE u.user_id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.Id, &video.Title, &video.Description, &video.CreatedAt)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
