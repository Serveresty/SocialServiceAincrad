package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func GetVideoById(vid string) (models.Video, error) {
	row := database.DB.QueryRow(context.Background(), `SELECT video_id, title, description, created_at, filename FROM "videos" WHERE video_id = $1`, vid)
	var video models.Video

	err := row.Scan(&video.Id, &video.Title, &video.Description, &video.CreatedAt, &video.Filename)
	if err != nil {
		return models.Video{}, err
	}

	return video, nil
}
