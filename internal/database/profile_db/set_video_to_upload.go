package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func SetVideoToUpload(video models.Video, filename string) (int, error) {
	var id int
	err := database.DB.QueryRow(context.Background(), `INSERT INTO "videos" (title, description, filename, created_at) VALUES($1, $2, $3, $4) RETURNING video_id`, video.Title, video.Description, filename, video.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
